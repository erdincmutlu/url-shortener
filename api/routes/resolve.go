package routes

import (
	"net/http"

	"github.com/erdincmutlu/url-shortener/api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short not found in the database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to db",
		})
	}

	// Increment the counter
	rInr := database.CreateClient(1)
	defer rInr.Close()
	_ = rInr.Incr(database.Ctx, "counter")

	// Redirect to original URL
	return c.Redirect(value, http.StatusMovedPermanently)
}
