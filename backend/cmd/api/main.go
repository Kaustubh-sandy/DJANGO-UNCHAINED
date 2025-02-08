package main

import (
	"log"
	"net/http"

	// "github.com/ABHINAVGARG05/DJANGO_UNCHAINED/pkg/db"
	"github.com/ABHINAVGARG05/DJANGO_UNCHAINED/pkg/db"
	"github.com/ABHINAVGARG05/DJANGO_UNCHAINED/pkg/models"
	"github.com/ABHINAVGARG05/DJANGO_UNCHAINED/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.DbConnection()

	routes.AuthRoutes(app)
	routes.SetupRoutes(app)
	// db.GetCollection("users")
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&models.Response{
			Status:  "success",
			Message: "pong",
		})
	})

	app.Listen(":8080")

}
