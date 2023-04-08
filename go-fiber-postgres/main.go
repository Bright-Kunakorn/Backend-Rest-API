package main

import (
	"log"
	"net/http"
	"os"

	"github.com/akhil/go-fiber-postgres/models"
	"github.com/akhil/go-fiber-postgres/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Product struct {
	SKU          string   `json:"sku"`
	Title        string   `json:"title"`
	Price        float64  `json:"price"`
	UpdateStatus bool     `json:"updateStatus"`
	BlueFlag     bool     `json:"blueFlag"`
	IsVat        bool     `json:"isVat"`
	IsFixPrice   bool     `json:"isFixPrice"`
	Promotion    []string `json:"promotion"`
}

func (r *Repository) GetProduct(context *fiber.Ctx) error {
	productModels := &[]models.Product{}
	err := r.DB.Find(productModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get product"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "product fethed successfuly",
		"data":    productModels,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/get", r.GetProduct)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_Port"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not lead the database")
	}
	err = models.MigrateProduct(db)
	if err != nil {
		log.Fatal("could not migrate database")
	}
	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
