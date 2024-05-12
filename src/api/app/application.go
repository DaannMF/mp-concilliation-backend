/*
Package app implement startup logic to run the application and set environment variables.
*/
package app

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/proethics/mp-conciliation/src/api/config/config_service"
	"github.com/proethics/mp-conciliation/src/api/config/database"
	"github.com/proethics/mp-conciliation/src/api/core/errors"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/dependencies"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
)

const (
	defaultPort = ":8080"
)

func Start() {
	config_service.SetupConfig()
	logger.SetupLogger()
	ctx := context.Background()
	ctx = context.WithValue(ctx, logger.MpConciliationKey{}, "mp-conciliation")

	logEnvironment(ctx)

	router := gin.Default()

	dependencies := dependencies.StartConnection{StoreConnection: new(database.GormConnection)}

	handlers := dependencies.Start()

	configureMappings(router, handlers)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := router.Run(port)
	if err != nil {
		logger.Error(ctx, errors.ErrorRunningApplication.GetMessage(), logger.Tags{})
		panic(err)
	}
}

func logEnvironment(ctx context.Context) {
	tags := logger.Tags{}
	logger.Info(ctx, "Starting Mercado Pago Conciliation!", tags)
	logger.Info(ctx, fmt.Sprintf("GO_ENVIRONMENT: %s", os.Getenv("GO_ENVIRONMENT")), tags)
	logger.Info(ctx, fmt.Sprintf("SCOPE: %s", os.Getenv("SCOPE")), tags)
}
