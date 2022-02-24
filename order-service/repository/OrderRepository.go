package repository

import (
	"github.com/ogikuzma/k8s/order-service/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	RWDatabase *gorm.DB
	RDatabase *gorm.DB
}

func (repo *OrderRepository) CreateOrder(order *model.Order) error {
	result := repo.RWDatabase.Create(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *OrderRepository) UpdateOrder(orderId uint, status int) error {
	result := repo.RWDatabase.Model(&model.Order{}).Where("id = ?", orderId).Update("order_status", status)
	if result.Error != nil {
		return result.Error
	}
	
	return nil
}
