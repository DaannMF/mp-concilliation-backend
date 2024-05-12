package consumer

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/proethics/mp-conciliation/src/api/core/errors"
	processPaymentsNews "github.com/proethics/mp-conciliation/src/api/core/usecases/process_payments_news"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
)

const (
	ParseBaseNumber = 10
	BitSize         = 64
)

type PaymentsNews struct {
	ProcessPaymentsNewsUseCase processPaymentsNews.UseCase
}

func (handler *PaymentsNews) Handle(c *gin.Context) {
	handler.handle(c)
}

func (handler *PaymentsNews) handle(c *gin.Context) {
	ctx := context.WithValue(c, logger.MpConciliationKey{}, "process_payments_news_handler")

	param := c.Query("id")
	payment_id, err := strconv.ParseInt(param, ParseBaseNumber, BitSize)
	if err != nil {
		logger.Error(ctx, errors.ErrorInvalidID.GetMessageWithParams(errors.Parameters{"id": c.Param("id")}), logger.Tags{})
		c.JSON(http.StatusOK, new(interface{}))
		return
	}

	logger.Info(ctx, fmt.Sprintf("%s%d", "Payment notification with ID : ", payment_id), logger.Tags{})

	err = handler.ProcessPaymentsNewsUseCase.Execute(ctx, payment_id)
	if err != nil {
		c.JSON(http.StatusOK, new(interface{}))
		return
	}

	c.JSON(http.StatusOK, new(interface{}))
}
