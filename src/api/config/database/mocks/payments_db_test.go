package mocks

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettingNewMockedDB(t *testing.T) {
	// given
	_ = os.Setenv("SCOPE", "local")

	// when
	mockDB, db := NewDB()

	// then
	assert.NotNil(t, mockDB)
	assert.NotNil(t, db)
}
