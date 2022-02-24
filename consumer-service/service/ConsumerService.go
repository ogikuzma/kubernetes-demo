package service

import (
	"strconv"

	"github.com/ogikuzma/k8s/model"
	"github.com/ogikuzma/k8s/repository"
)

type ConsumerService struct {
	Repo *repository.ConsumerRepository
}

func (service *ConsumerService) CreateConsumer(consumer *model.Consumer) error {
	err := service.Repo.CreateConsumer(consumer)
	if err != nil {
		return err
	}

	return nil
}

func (service *ConsumerService) UserExists(consumerId string) (bool, error) {
	var id uint64
	id, error := strconv.ParseUint(consumerId, 10, 32)
	if error != nil {
		return false, error
	}

	exists := service.Repo.ConsumerExists(uint(id))
	return exists, nil
}
