package entities

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	UserName  string    `gorm:"type:varchar(50);unique"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"column:date_created;type:datetime;"`
	UpdatedAt time.Time `gorm:"column:date_updated;type:datetime;"`
}
