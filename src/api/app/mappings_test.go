package app

import (
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/proethics/mp-conciliation/src/api/config/config_service"
	"github.com/proethics/mp-conciliation/src/api/config/database"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/dependencies"
	"github.com/proethics/mp-conciliation/src/api/utils"
	"github.com/stretchr/testify/assert"
)

func TestUrlMappings(t *testing.T) {
	// Given
	_ = os.Setenv("SCOPE", "test")
	config_service.SetupConfig()
	router := utils.GetTestRouter()
	dependencies := dependencies.StartConnection{StoreConnection: new(database.GormConnection)}

	handlers := dependencies.Start()

	// When
	configureMappings(router, handlers)

	var consumers []gin.RouteInfo

	routes := router.Routes()
	for _, r := range routes {
		if strings.Contains(r.Path, "/consumers/") {
			consumers = append(consumers, r)
			continue
		}
	}

	// Then
	assert.NotNil(t, router)
	assert.EqualValues(t, 1, len(consumers))
}
