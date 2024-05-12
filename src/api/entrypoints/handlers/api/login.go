package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/proethics/mp-conciliation/src/api/core/contracts/auth"
	"github.com/proethics/mp-conciliation/src/api/core/usecases/login"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
)

type Login struct {
	LoginUseCase login.UseCase
}

func (handler *Login) Handle(c *gin.Context) {
	handler.handle(c)
}

func (handler *Login) handle(c *gin.Context) {
	ctx := context.WithValue(c, logger.MpConciliationKey{}, "login_handler")

	var authInput auth.Request
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := handler.LoginUseCase.Execute(ctx, authInput)
	if err != nil {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
