package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ogikuzma/k8s/order-service/handler"
	"github.com/ogikuzma/k8s/order-service/model"
	"github.com/ogikuzma/k8s/order-service/repository"
	"github.com/ogikuzma/k8s/order-service/service"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB(host string) *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbport)

	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			time.Sleep(5 * time.Second)
			log.Println(err)
		} else {
			log.Println("Connected to db...")
			break
		}
	}

	if err != nil{
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.OrderItem{})

	/*Loading test data*/
	orders := []model.Order{
		{ConsumerID: 1, OrderStatus: model.PENDING, OrderItems: []model.OrderItem{
			{MenuItemName: "jaja", Quantity: 6},
			{MenuItemName: "palacinke", Quantity: 1},
		}},
		{ConsumerID: 2, OrderStatus: model.ACCEPTED, OrderItems: []model.OrderItem{
			{MenuItemName: "carbonara", Quantity: 2},
			{MenuItemName: "mleko", Quantity: 3},
		}},
	}
	for i := range orders {
		db.Create(&orders[i])
	}

	return db
}

func initRepo(rwDatabase *gorm.DB, rDatabase *gorm.DB) *repository.OrderRepository {
	return &repository.OrderRepository{RWDatabase: rwDatabase, RDatabase: rDatabase}
}

func initServices(repo *repository.OrderRepository) *service.OrderService {
	return &service.OrderService{Repo: repo}
}

func initHandler(service *service.OrderService) *handler.OrderHandler {
	return &handler.OrderHandler{Service: service}
}
func handleFunc(handler *handler.OrderHandler) {
	app := fiber.New()

	api := app.Group("/api/order")

	api.Get("/hello", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	api.Post("/", handler.CreateOrder)
	api.Put("/:orderId/:status", handler.UpdateStatus)

	log.Println("server running...")
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func main() {
	rwDatabase := initDB(os.Getenv("DB_PRIMARY"))
	rDatabase := initDB(os.Getenv("DB_REPLICA"))
	repo := initRepo(rwDatabase, rDatabase)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}
