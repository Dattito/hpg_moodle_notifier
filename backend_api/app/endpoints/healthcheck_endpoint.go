package endpoints

import "github.com/gofiber/fiber/v2"

func GetHeartbeat(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
