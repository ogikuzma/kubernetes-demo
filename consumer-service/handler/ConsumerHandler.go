package handler

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/ogikuzma/k8s/model"
	"github.com/ogikuzma/k8s/service"
)

type ConsumerHandler struct {
	Service *service.ConsumerService
}

func (handler *ConsumerHandler) CreateConsumer(c *fiber.Ctx) error {
	log.Println("creating consumer...")

	var consumer model.Consumer
	err := c.BodyParser(&consumer)
	if err != nil {
		log.Println("cannot parse customer json")
		return err
	}

	err = handler.Service.CreateConsumer(&consumer)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{
		"success": "ok",
	})
}

func (handler *ConsumerHandler) Verify(c *fiber.Ctx) error {
	log.Println("verifying if consumer exists... ")

	consumerId := c.Params("consumerId")
	if consumerId == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "consumerId is missing!",
		})
	}

	exists, err := handler.Service.UserExists(consumerId)
	if err != nil {
		return err
	}

	if exists {
		return c.SendStatus(200)
	} else {
		return c.SendStatus(404)
	}
}
