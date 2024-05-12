package entities

import "github.com/shopspring/decimal"

type TransactionDetails struct {
	ID                       int64            `gorm:"primaryKey;autoIncrement"`
	PaymentID                int64            `gorm:"type:bigint(20);not null"`
	NetReceivedAmount        *decimal.Decimal `gorm:"type:decimal(11,2)"`
	TotalPaidAmount          *decimal.Decimal `gorm:"type:decimal(11,2)"`
	OverpaidAmount           *decimal.Decimal `gorm:"type:decimal(11,2)"`
	ExternalResourceURL      *string          `gorm:"type:varchar(100)"`
	InstallmentAmount        *decimal.Decimal `gorm:"type:decimal(11,2)"`
	FinancialInstitution     *string          `gorm:"type:varchar(100)"`
	PaymentMethodReferenceID *string          `gorm:"type:varchar(100)"`
	PayableDeferralPeriod    *string          `gorm:"type:varchar(100)"`
	AcquirerReference        *string          `gorm:"type:varchar(100)"`
}
