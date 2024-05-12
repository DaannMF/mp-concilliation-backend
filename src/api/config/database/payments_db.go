package database

import (
	"context"
	"time"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/errors"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	maxLifetime = time.Minute * 30
	maxIdleConn = 5
	maxOpenConn = 5
)

type Connection interface {
	Connect() (client *gorm.DB, err error)
}

type GormConnection struct{}

func (connection GormConnection) Connect() (client *gorm.DB, err error) {
	connectionData := GetConnectionDataBase()
	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 gormLogger.Discard,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	client, err = gorm.Open(connectionData.DialectConnect(GetConnectionString(connectionData)), &config)

	if err != nil {
		logger.Error(context.Background(), errors.ErrorDataBaseConnection.GetMessage(), logger.Tags{})
		panic(err)
	}

	sqlDB, err := client.DB()
	if err != nil {
		logger.Error(context.Background(), errors.ErrorDataBaseConnection.GetMessage(), logger.Tags{})
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(maxLifetime)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)

	err = client.AutoMigrate(
		&entities.Payment{},
		&entities.PaymentPayer{},
		&entities.TransactionDetails{},
		&entities.User{},
	)

	if err != nil {
		logger.Error(context.Background(), errors.ErrorDataBaseMigration.GetMessage(), logger.Tags{})
		panic(err)
	}

	logger.Info(context.Background(), errors.InfoDataBaseConnection.GetMessageWithParams(errors.Parameters{
		"database": connectionData.Schema,
	}), logger.Tags{})

	return client, nil
}
