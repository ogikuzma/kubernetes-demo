package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ogikuzma/k8s/order-service/model"
	"github.com/ogikuzma/k8s/order-service/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func (service *OrderService) CreateOrder(order *model.Order) error {
	err_consumer := verifyConsumer(os.Getenv("CONSUMER_SERVICE_DOMAIN"), os.Getenv("CONSUMER_SERVICE_PORT"), order.ConsumerID)
	if err_consumer != nil{
		return err_consumer
	}

	order.OrderStatus = model.PENDING
	err := service.Repo.CreateOrder(order)
	if err != nil {
		return err
	}
	
	return nil
}

func verifyConsumer(domain, port string, id uint) error {
	var consumerId string = strconv.FormatUint(uint64(id), 10)
	url := fmt.Sprintf("http://%s:%s/api/consumer/verify/%s", domain, port, consumerId)
	log.Println("Verifying for url %s\n", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode == 404 {
		log.Printf("Verification failed for domain %s and id %s\n", domain, consumerId)
		return fmt.Errorf(fmt.Sprintf("verification failed for domain %s and id %s", domain, consumerId))
	}
	return nil
}

func (service *OrderService) ChangeStatus(orderId string, status string) error {
	var id uint64
	id, error := strconv.ParseUint(orderId, 10, 32)
	if error != nil {
		return error
	}

	var orderStatus int
	switch status {
	case "pending":
		orderStatus = model.PENDING
	case "accepted":
		orderStatus = model.ACCEPTED
	case "rejected":
		orderStatus = model.REJECTED
	}
	service.Repo.UpdateOrder(uint(id), orderStatus)
	return nil
}
