package model

type Consumer struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
}
