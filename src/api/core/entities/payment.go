package entities

import (
	"time"

	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
	"github.com/shopspring/decimal"
)

type Payment struct {
	ID                        int64                         `gorm:"primaryKey;autoIncrement"`
	DateCreated               *time.Time                    `gorm:"column:date_created;type:datetime;"`
	DateApproved              *time.Time                    `gorm:"column:date_approved;type:datetime;"`
	DateLastUpdated           *time.Time                    `gorm:"column:date_last_updated;type:datetime;"`
	DateOfExpiration          *time.Time                    `gorm:"column:date_expiration;type:datetime;"`
	MoneyReleaseDate          *time.Time                    `gorm:"column:money_release_date;type:datetime;"`
	OperationType             constants.OperationType       `gorm:"type:varchar(25);not null"`
	IssuerID                  string                        `gorm:"type:varchar(50);not null"`
	PaymentMethodID           constants.PaymentMethodID     `gorm:"type:varchar(15);not null"`
	PaymentTypeID             constants.PaymentTypeID       `gorm:"type:varchar(20);not null"`
	Status                    constants.PaymentStatus       `gorm:"type:varchar(15);not null"`
	StatusDetail              constants.PaymentStatusDetail `gorm:"type:varchar(50);not null"`
	CurrencyID                constants.Currency            `gorm:"type:varchar(5);not null"`
	Description               *string                       `gorm:"type:varchar(150)"`
	LiveMode                  bool                          `gorm:"type:tinyint(1);default:false"`
	AuthorizationCode         *string                       `gorm:"type:varchar(6)"`
	Payer                     PaymentPayer                  `gorm:"foreignkey:PaymentID;association_autoupdate:false"`
	ExternalReference         *string                       `gorm:"type:varchar(100)"`
	TransactionAmount         decimal.Decimal               `gorm:"type:decimal(11,2);not null"`
	TransactionAmountRefunded *decimal.NullDecimal          `gorm:"type:decimal(11,2);"`
	CouponAmount              *decimal.NullDecimal          `gorm:"type:decimal(11,2);"`
	DeductionSchema           *string                       `gorm:"type:varchar(20)"`
	TransactionDetails        TransactionDetails            `gorm:"foreignkey:PaymentID;association_autoupdate:false"`
	Captured                  bool                          `gorm:"type:tinyint(1);default:false"`
	BinaryMode                bool                          `gorm:"type:tinyint(1);default:false"`
	CallForAuthorizeID        *string                       `gorm:"type:varchar(20)"`
	StatementDescriptor       *string                       `gorm:"type:varchar(20)"`
	Installments              int64                         `gorm:"type:bigint(20);not null"`
	ConcilliedStatus          constants.ConcilliedStatus    `gorm:"type:varchar(10)"`
	ConcilliedUser            *string                       `gorm:"type:varchar(50)"`
	ConcilliedDate            *time.Time                    `gorm:"column:concillied_date;type:datetime;"`
}
