package entities

import (
	"time"

	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
)

type User struct {
	ID        int64              `gorm:"primaryKey;autoIncrement"`
	UserName  string             `gorm:"type:varchar(50);unique"`
	Password  string             `gorm:"type:varchar(100);not null"`
	UserRole  constants.UserRole `gorm:"column:role;type:varchar(10);not null"`
	CreatedAt time.Time          `gorm:"column:date_created;type:datetime;"`
	UpdatedAt time.Time          `gorm:"column:date_updated;type:datetime;"`
}
