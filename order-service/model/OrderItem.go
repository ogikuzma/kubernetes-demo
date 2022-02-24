package model

type OrderItem struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	MenuItemName string `json:"item_name"`
	Quantity     int    `json:"quantity" gorm:"not null"`
	OrderID      uint
}
