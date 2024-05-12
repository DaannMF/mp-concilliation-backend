package utils

import (
	"github.com/gin-gonic/gin"
)

const BaseURL = "http://127.0.0.1:8080"

func GetTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	return router
}
