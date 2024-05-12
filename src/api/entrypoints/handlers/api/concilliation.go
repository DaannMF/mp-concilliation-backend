package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/errors"
	"github.com/proethics/mp-conciliation/src/api/core/usecases/concilliation"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
)

const (
	ParseBaseNumber = 10
	BitSize         = 64
)

type Concilliation struct {
	ConcilliationUseCase concilliation.UseCase
}

func (handler *Concilliation) Handle(c *gin.Context) {
	handler.handle(c)
}

func (handler *Concilliation) handle(c *gin.Context) {
	ctx := context.WithValue(c, logger.MpConciliationKey{}, "concilliation_handler")

	param := c.Param("payment_id")
	value, exists := c.Get("currentUser")
	if !exists {
		logger.Error(ctx, errors.ErrorUserNotPresent.GetMessage(), logger.Tags{})
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	user, exists := value.(entities.User)
	if !exists {
		logger.Error(ctx, errors.ErrorUserNotPresent.GetMessage(), logger.Tags{})
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	payment_id, err := strconv.ParseInt(param, ParseBaseNumber, BitSize)
	if err != nil {
		logger.Error(ctx, errors.ErrorInvalidID.GetMessage(), logger.Tags{"payment_id": param})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = handler.ConcilliationUseCase.Execute(ctx, payment_id, user.UserName)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, new(interface{}))
}
