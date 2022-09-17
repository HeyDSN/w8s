package models

import "time"

type Order struct {
	OrderId      uint64    `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `gorm:"foreignKey:OrderID" json:"items"`
}

type Item struct {
	ItemId      uint64 `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
	OrderID     uint64 `json:"order_id"`
}
