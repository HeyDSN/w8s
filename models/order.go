package models

import "time"

type Order struct {
	OrderId      uint64 `gorm:"primaryKey"`
	CustomerName string
	OrderedAt    time.Time
	Items        []Item `gorm:"foreignKey:OrderID"`
}

type Item struct {
	ItemId      uint64 `gorm:"primaryKey"`
	ItemCode    string
	Description string
	Quantity    uint64
	OrderID     uint64
}
