package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/proethics/mp-conciliation/src/api/core/contracts/search"
	"github.com/proethics/mp-conciliation/src/api/core/providers"
	"github.com/proethics/mp-conciliation/src/api/infrastructure/logger"
)

type Search struct {
	PaymentsProvider providers.Payments
}

func (handler *Search) Handle(c *gin.Context) {
	handler.handle(c)
}

func (handler *Search) handle(c *gin.Context) {
	ctx := context.WithValue(c, logger.MpConciliationKey{}, "get_payments_handler")

	payments, err := handler.PaymentsProvider.GetPendingPayments(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, new(interface{}))
		return
	}

	c.JSON(http.StatusOK, search.NewSearchResponse(payments))
}
