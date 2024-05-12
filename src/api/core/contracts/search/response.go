package search

import (
	"time"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
	"github.com/shopspring/decimal"
)

type SearchResponse struct {
	PaymentID            int64                        `json:"payment_id"`
	MoneyReleaseDate     *time.Time                   `json:"money_release_date"`
	PaymentStatus        constants.PaymentStatus      `json:"payment_status"`
	Email                string                       `json:"payer_email"`
	IdentificationType   constants.IdentificationType `json:"payer_identification_type"`
	IdentificationNumber string                       `json:"payer_identification_number"`
	CurrencyID           constants.Currency           `json:"currency_id"`
	TransactionAmount    decimal.Decimal              `json:"transaction_amount"`
}

func NewSearchResponse(payments []entities.Payment) []SearchResponse {
	response := make([]SearchResponse, len(payments))

	for i, payment := range payments {
		res := SearchResponse{
			PaymentID:            payment.ID,
			MoneyReleaseDate:     payment.MoneyReleaseDate,
			PaymentStatus:        payment.Status,
			Email:                payment.Payer.Email,
			IdentificationType:   payment.Payer.Type,
			IdentificationNumber: payment.Payer.Number,
			CurrencyID:           payment.CurrencyID,
			TransactionAmount:    payment.TransactionAmount,
		}

		response[i] = res
	}

	return response
}
