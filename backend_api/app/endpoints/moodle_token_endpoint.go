package endpoints

import (
	"log"

	"github.com/Dattito/HMN_backend_api/app/utils"
	"github.com/gofiber/fiber/v2"
)

type MoodleVerificationRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func PostMoodleToken(c *fiber.Ctx) error {
	mvr := &MoodleVerificationRequest{}
	if err := c.BodyParser(mvr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := utils.NewValidator().Struct(mvr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": utils.ValidatorErrors(err),
		})
	}

	token, err := utils.GetMoodleToken(mvr.Username, mvr.Password)
	if err != nil {

		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Etwas ist schiefgelaufen.",
		})
	}

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"msg": "Falsche Anmeldedaten.",
		})
	}

	return c.JSON(&fiber.Map{
		"token": token,
	})
}
