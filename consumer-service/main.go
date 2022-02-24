package main

import (
	"fmt"
	"log"
	"os"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/ogikuzma/k8s/handler"
	"github.com/ogikuzma/k8s/model"
	"github.com/ogikuzma/k8s/repository"
	"github.com/ogikuzma/k8s/service"
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

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Consumer{})

	/*Loading test data*/
	consumers := []model.Consumer{
		{Email: "petar.petrovic@mail.cc", Password: "petar", Name: "petar", Surname: "petrovic"},
		{Email: "ivan.ivanovic@example.cc", Password: "ivan", Name: "ivan", Surname: "ivanovic"},
	}
	for _, consumer := range consumers {
		db.Create(&consumer)
	}
	return db
}

func initRepo(rwDatabase *gorm.DB, rDatabase *gorm.DB) *repository.ConsumerRepository {
	return &repository.ConsumerRepository{RWDatabase: rwDatabase, RDatabase: rDatabase}
}

func initServices(repo *repository.ConsumerRepository) *service.ConsumerService {
	return &service.ConsumerService{Repo: repo}
}

func initHandler(service *service.ConsumerService) *handler.ConsumerHandler {
	return &handler.ConsumerHandler{Service: service}
}
func handleFunc(handler *handler.ConsumerHandler) {
	app := fiber.New()

	api := app.Group("/api/consumer")

	api.Get("/helloo", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	api.Get("/crash", func(c *fiber.Ctx) error {
		go func() {
			time.Sleep(1 * time.Second)
			os.Exit(1)
		}()
		return c.SendString("crashing...")
	})

	api.Post("/", handler.CreateConsumer)
	api.Get("/verify/:consumerId", handler.Verify)

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
