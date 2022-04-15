package controllers

import (
	"example.com/web-service-gin/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

func HandlerHealth(c *fiber.Ctx) error {
	repositories.POSTHealth()
	return c.SendStatus(fiber.StatusOK)
}

func HandlerDeleteRestaurant(c *fiber.Ctx) error {
	restaurantId := c.Params("restaurantId")
	id := repositories.DeleteRestaurant(restaurantId)
	return c.SendString(id)
}

func HandlerListRestaurant(c *fiber.Ctx) error {
	term := c.Query("name")
	if term != "" {
		restaurants := repositories.FindRestaurant(term)
		return c.Send(*restaurants)
	}
	restaurants := repositories.GETListAllRestaurants()
	return c.Send(*restaurants)
}
