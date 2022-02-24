package model

type Order struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	ConsumerID  uint `json:"consumer_id"`
	OrderStatus int
	OrderItems  []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}
