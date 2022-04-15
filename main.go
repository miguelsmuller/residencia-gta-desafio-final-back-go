package main

import (
	"example.com/web-service-gin/internal/controllers"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	app.Post("/health", controllers.HandlerHealth)
	app.Delete("/restaurants/:restaurantId", controllers.HandlerDeleteRestaurant)
	app.Get("/restaurants", controllers.HandlerListRestaurant)

	app.Listen(":3009")
}
