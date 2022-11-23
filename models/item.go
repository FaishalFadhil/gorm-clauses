package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	Id          uint   `gorm:"PRIMARY_KEY"`
	ItemCode    string `gorm:"column:item_code;varchar(100)"`
	Description string `gorm:"column:description;varchar(100)"`
	Quantity    int32  `gorm:"column:quantity;integer(11)"`
	OrderId     uint   `gorm:"column:order_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
