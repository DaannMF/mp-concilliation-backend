package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/proethics/mp-conciliation/src/api/core/contracts/search"
	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
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
	value, present := c.GetQuery("status")
	if !present {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid query param, status is required"})
		return
	}

	status := constants.ParseConcilliedStatus(value)
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid status, possible values : " + strings.Join(status.GetValues(), ","),
		})
		return
	}

	payments, err := handler.PaymentsProvider.GetPaymentsByStatus(ctx, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, new(interface{}))
		return
	}

	c.JSON(http.StatusOK, search.NewSearchResponse(payments))
}
