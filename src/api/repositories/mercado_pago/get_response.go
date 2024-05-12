package mercado_pago

import (
	"time"

	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
	"github.com/shopspring/decimal"
)

type response struct {
	ID                        int64                         `json:"id"`
	Number                    string                        `json:"number"`
	DateCreated               *time.Time                    `json:"date_created"`
	DateApproved              *time.Time                    `json:"date_approved"`
	DateLastUpdated           *time.Time                    `json:"date_last_updated"`
	DateOfExpiration          *time.Time                    `json:"date_of_expiration"`
	MoneyReleaseDate          *time.Time                    `json:"money_release_date"`
	OperationType             constants.OperationType       `json:"operation_type"`
	IssuerID                  string                        `json:"issuer_id"`
	PaymentMethodID           constants.PaymentMethodID     `json:"payment_method_id"`
	PaymentTypeID             constants.PaymentTypeID       `json:"payment_type_id"`
	Status                    constants.PaymentStatus       `json:"status"`
	StatusDetail              constants.PaymentStatusDetail `json:"status_detail"`
	CurrencyID                constants.Currency            `json:"currency_id"`
	Description               string                        `json:"description"`
	LiveMode                  bool                          `json:"live_mode"`
	AuthorizationCode         string                        `json:"authorization_code"`
	MoneyReleaseSchema        string                        `json:"money_release_schema"`
	CounterCurrency           string                        `json:"counter_currency"`
	CollectorID               int64                         `json:"collector_id"`
	Payer                     Payer                         `json:"payer"`
	ExternalReference         string                        `json:"external_reference"`
	TransactionAmount         decimal.Decimal               `json:"transaction_amount"`
	TransactionAmountRefunded *decimal.NullDecimal          `json:"transaction_amount_refunded"`
	CouponAmount              *decimal.NullDecimal          `json:"coupon_amount"`
	DifferentialPricingID     string                        `json:"differential_pricing_id"`
	DeductionSchema           string                        `json:"deduction_schema"`
	TransactionDetails        TransactionDetails            `json:"transaction_details"`
	AcquirerReference         string                        `json:"acquirer_reference"`
	Captured                  bool                          `json:"captured"`
	BinaryMode                bool                          `json:"binary_mode"`
	CallForAuthorizeID        string                        `json:"call_for_authorize_id"`
	StatementDescriptor       string                        `json:"statement_descriptor"`
	Installments              int64                         `json:"installments"`
}

type Payer struct {
	ID             string              `json:"id"`
	Email          string              `json:"email"`
	Identification Identification      `json:"identification"`
	Type           constants.PayerType `json:"type"`
}

type Identification struct {
	Type   constants.IdentificationType `json:"type"`
	Number string                       `json:"number"`
}

type TransactionDetails struct {
	NetReceivedAmount        *decimal.Decimal `json:"net_received_amount"`
	TotalPaidAmount          *decimal.Decimal `json:"total_paid_amount"`
	OverpaidAmount           *decimal.Decimal `json:"overpaid_amount"`
	ExternalResourceURL      *string          `json:"external_resource_url"`
	InstallmentAmount        *decimal.Decimal `json:"installment_amount"`
	FinancialInstitution     *string          `json:"financial_institution"`
	PaymentMethodReferenceID *string          `json:"payment_method_reference_id"`
	PayableDeferralPeriod    *string          `json:"payable_deferral_period"`
	AcquirerReference        *string          `json:"acquirer_reference"`
}

func (response response) GetEntity() entities.Payment {
	return entities.Payment{
		ID:                response.ID,
		DateCreated:       response.DateCreated,
		DateApproved:      response.DateApproved,
		DateLastUpdated:   response.DateLastUpdated,
		DateOfExpiration:  response.DateOfExpiration,
		MoneyReleaseDate:  response.MoneyReleaseDate,
		OperationType:     response.OperationType,
		IssuerID:          response.IssuerID,
		PaymentMethodID:   response.PaymentMethodID,
		PaymentTypeID:     response.PaymentTypeID,
		Status:            response.Status,
		StatusDetail:      response.StatusDetail,
		CurrencyID:        response.CurrencyID,
		Description:       &response.Description,
		LiveMode:          response.LiveMode,
		AuthorizationCode: &response.AuthorizationCode,
		Payer: entities.PaymentPayer{
			PayerID:   response.Payer.ID,
			PaymentID: response.ID,
			Email:     response.Payer.Email,
			Type:      response.Payer.Identification.Type,
			Number:    response.Payer.Identification.Number,
			PayerType: response.Payer.Type,
		},
		ExternalReference:         &response.ExternalReference,
		TransactionAmount:         response.TransactionAmount,
		TransactionAmountRefunded: response.TransactionAmountRefunded,
		CouponAmount:              response.CouponAmount,
		DeductionSchema:           &response.DeductionSchema,
		TransactionDetails: entities.TransactionDetails{
			PaymentID:                response.ID,
			NetReceivedAmount:        response.TransactionDetails.NetReceivedAmount,
			TotalPaidAmount:          response.TransactionDetails.TotalPaidAmount,
			OverpaidAmount:           response.TransactionDetails.OverpaidAmount,
			ExternalResourceURL:      response.TransactionDetails.ExternalResourceURL,
			InstallmentAmount:        response.TransactionDetails.InstallmentAmount,
			FinancialInstitution:     response.TransactionDetails.FinancialInstitution,
			PaymentMethodReferenceID: response.TransactionDetails.PaymentMethodReferenceID,
			PayableDeferralPeriod:    response.TransactionDetails.PayableDeferralPeriod,
			AcquirerReference:        response.TransactionDetails.AcquirerReference,
		},
		Captured:            response.Captured,
		BinaryMode:          response.BinaryMode,
		CallForAuthorizeID:  &response.CallForAuthorizeID,
		StatementDescriptor: &response.StatementDescriptor,
		Installments:        response.Installments,
	}
}
