/*
Package database implement the logic to set database configuration by environment
*/
package database

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	mysqlDialect   = "mysql"
	sqlLiteDialect = "sqlite3"
)

type ConnectionData struct {
	Host           string
	Port           string
	Schema         string
	Username       string
	Password       string
	Dialect        string
	DialectConnect func(dns string) gorm.Dialector
}

func GetConnectionDataBase() *ConnectionData {
	scope := os.Getenv("SCOPE")

	connectionData := ConnectionData{}

	if strings.HasSuffix(scope, "test") {
		return connectionData.setupTestConnectionData()
	}

	if strings.HasSuffix(scope, "master") {
		return connectionData.setupMasterConnectionData()
	}

	connectionData.Host = "localhost"
	connectionData.Port = "3306"
	connectionData.Schema = "payments"
	connectionData.Username = "root"
	connectionData.Password = ""
	connectionData.Dialect = mysqlDialect
	connectionData.DialectConnect = mySQLConnect
	return &connectionData
}

func (cd *ConnectionData) setupMasterConnectionData() *ConnectionData {
	cd.Host = os.Getenv("MYSQLHOST")
	cd.Port = os.Getenv("MYSQLPORT")
	cd.Password = os.Getenv("MYSQLPASSWORD")
	cd.Username = os.Getenv("MYSQLUSER")
	cd.Schema = "payments"
	cd.Dialect = mysqlDialect
	cd.DialectConnect = mySQLConnect
	return cd
}

func (cd *ConnectionData) setupTestConnectionData() *ConnectionData {
	cd.Dialect = sqlLiteDialect
	cd.DialectConnect = sqlLiteConnect
	cd.Host = "file::memory:?cache=shared"
	cd.Schema = "credlines"
	return cd
}

func GetConnectionString(cd *ConnectionData) string {
	if cd.Dialect == sqlLiteDialect {
		return cd.Host
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		cd.Username, cd.Password, cd.Host, cd.Port, cd.Schema)
}

func mySQLConnect(dns string) gorm.Dialector {
	return mysql.Open(dns)
}

func sqlLiteConnect(dns string) gorm.Dialector {
	return sqlite.Open(dns)
}
