package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id           uint   `gorm:"PRIMARY_KEY"`
	CustomerName string `gorm:"column:customer_name;varchar(100)"`
	Items        []Item `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE"`
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
