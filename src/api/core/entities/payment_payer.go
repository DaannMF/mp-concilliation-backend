package entities

import "github.com/proethics/mp-conciliation/src/api/core/entities/constants"

type PaymentPayer struct {
	ID        int64                        `gorm:"primaryKey;autoIncrement"`
	PayerID   string                       `gorm:"type:varchar(64)"`
	PaymentID int64                        `gorm:"type:bigint(20);not null"`
	Email     string                       `gorm:"type:varchar(64)"`
	Type      constants.IdentificationType `gorm:"type:varchar(10)"`
	Number    string                       `gorm:"type:varchar(64)"`
	PayerType constants.PayerType          `gorm:"type:varchar(15)"`
}
