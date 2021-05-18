package endpoints

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Dattito/HMN_backend_api/app/models"
	"github.com/Dattito/HMN_backend_api/app/utils"
	"github.com/Dattito/HMN_backend_api/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateSignalVerification godoc
// @Summary Verify PhoneNumber in Signal
// @Description get ID of the verification-row to check later if code is correct
// @Accept  json
// @Produce  json
// @Param phone_number form string true "Phone Number of user (in e164 format!)"
// @Param moodle_token path string true "Moodle Token the user wants to register"
// @Success 200 {object} CreateSignalVerificationResponse
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /v1/signalVerification [post]
func CreateSignalVerification(c *fiber.Ctx) error {
	sv := &models.SignalVerification{}

	if err := c.BodyParser(sv); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	sv.ID = uuid.New()
	sv.CreatedAt = time.Now()
	sv.UpdatedAt = time.Now()
	sv.VerificationCode = utils.GenerateVerificationCode()
	sv.ValidUntil = time.Now().Add(time.Minute * 5)
	sv.MoodleToken = strings.ToLower(sv.MoodleToken)

	if err := utils.NewValidator().Struct(sv); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			//"msg": utils.ValidatorErrors(err),
			"msg": utils.ValidationErrorsToText(err),
		})
	}

	token_validation, err := utils.CheckToken(sv.MoodleToken)

	if err != nil {
		log.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Etwas ist schiefgelaufen.",
		})
	}

	if !token_validation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Ungültiger Token.",
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		log.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Etwas ist schiefgelaufen.",
		})
	}

	check_sb, _ := db.GetSignalVerificationByMoodleToken(sv.MoodleToken)

	if *check_sb != (models.SignalVerification{}) {

		// check if verification code is invalid...
		if check_sb.ValidUntil.Before(time.Now()) || check_sb.PhoneNumber != sv.PhoneNumber {
			db.DeleteSignalVerification(check_sb)
		} else {
			return c.JSON(fiber.Map{
				"id": check_sb.ID.String(),
			})
		}
	}

	if err := utils.SendSignalMessage(sv.PhoneNumber, "Dein Verifizierungs-Code ist "+strconv.Itoa(sv.VerificationCode)); err != nil {
		log.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Etwas ist schiefgelaufen.",
		})
	}

	if err := db.CreateSignalVerification(sv); err != nil {
		log.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Etwas ist schiefgelaufen.",
		})
	}

	return c.JSON(fiber.Map{
		"id": sv.ID.String(),
	})

}

type CreateSignalVerificationResponse struct {
	ERROR bool   `json:"error"`
	ID    string `json:"id"`
}

type HTTPError struct {
	MSG   string `json:"msg"`
	ERROR bool   `json:"error"`
}
