package repository

import (
	"github.com/ogikuzma/k8s/model"
	"gorm.io/gorm"
)

type ConsumerRepository struct {
	RWDatabase *gorm.DB
	RDatabase *gorm.DB
}

func (repo *ConsumerRepository) CreateConsumer(consumer *model.Consumer) error {
	result := repo.RWDatabase.Create(consumer)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *ConsumerRepository) ConsumerExists(consumerId uint) bool {
	var count int64
	repo.RDatabase.Where("id = ?", consumerId).Find(&model.Consumer{}).Count(&count)
	return count != 0
}
