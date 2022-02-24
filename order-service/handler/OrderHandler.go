package handler

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/ogikuzma/k8s/order-service/model"
	"github.com/ogikuzma/k8s/order-service/service"
)

type OrderHandler struct {
	Service *service.OrderService
}

func (handler *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	log.Println("creating order...")

	var order model.Order
	err := c.BodyParser(&order)
	if err != nil {
		log.Println("cannot parse order json")
		return err
	}

	err = handler.Service.CreateOrder(&order)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{
		"success": "ok",
	})
}

func (handler *OrderHandler) UpdateStatus(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	if orderId == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "orderId is missing!",
		})
	}

	status := c.Params("status")
	if status == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "status is missing!",
		})
	}

	err := handler.Service.ChangeStatus(orderId, status)
	if err != nil {
		return err
	}

	return c.SendStatus(200)
}
