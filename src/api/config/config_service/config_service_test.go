package config_service

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironmentSetupWhenIsTestScopeShouldSetRightVariables(t *testing.T) {
	// Given
	err := os.Setenv("SCOPE", "test")

	// When
	SetupConfig()
	configFileName := os.Getenv("configFileName")
	checksumEnabled := os.Getenv("checksumEnabled")
	isProdScope := os.Getenv("IS_PROD_SCOPE")

	t.Logf("configFileName: %s", configFileName)
	t.Logf("checksumEnabled; %s", checksumEnabled)
	t.Logf("isProdScope: %s", isProdScope)

	// Then
	assert.NoError(t, err)
	if isProdScope != "true" {
		assert.True(t, strings.HasSuffix(configFileName, "/src/api/config/config_service/latest/application.properties"))
	}
	assert.NotNil(t, configFileName)
	assert.NotNil(t, checksumEnabled)
	assert.Equal(t, "false", checksumEnabled)
	assert.NotNil(t, isProdScope)
	assert.Equal(t, "false", isProdScope)
}
