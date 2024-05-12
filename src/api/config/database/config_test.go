package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettingDataBaseConnectionWhenIsMasterShouldReturnIt(t *testing.T) {
	// given
	_ = os.Setenv("SCOPE", "master")

	// when
	connectionDB := GetConnectionDataBase()

	// then
	assert.Equal(t, "", connectionDB.Host)
	assert.Equal(t, "credlines", connectionDB.Schema)
	assert.Equal(t, "credlines_WPROD", connectionDB.Username)
	assert.Equal(t, "mysql", connectionDB.Dialect)
}

func TestGettingDataBaseConnectionWhenIsLocalShouldReturnIt(t *testing.T) {
	// given
	_ = os.Setenv("SCOPE", "")

	// when
	connectionDB := GetConnectionDataBase()

	// then
	assert.Equal(t, "localhost:3306", connectionDB.Host)
	assert.Equal(t, "credlines", connectionDB.Schema)
	assert.Equal(t, "root", connectionDB.Username)
	assert.Equal(t, "mysql", connectionDB.Dialect)
}

func TestGettingConnectionStringWhenScopeIsMasterShouldReturnIt(t *testing.T) {
	// given
	_ = os.Setenv("SCOPE", "master")

	// when
	connectionMaster := GetConnectionString(GetConnectionDataBase())

	// then
	assert.Equal(t, "credlines_WPROD:@tcp()/credlines?charset=utf8&parseTime=true", connectionMaster)
}
