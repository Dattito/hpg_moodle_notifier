package endpoints

import (
	"log"
	"strings"
	"time"

	"github.com/Dattito/HMN_backend_api/app/models"
	"github.com/Dattito/HMN_backend_api/app/utils"
	"github.com/Dattito/HMN_backend_api/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SignalVerificationRequest struct {
	ID               uuid.UUID `json:"id" validate:"required,uuid"`
	VerificationCode int       `json:"verificationCode" validate:"required,min=111111,max=999999"`
}

type responseError struct {
	Description *fiber.Map
	StatusCode  int
}

func verifiyRequest(db *database.Queries, c *fiber.Ctx, svr *SignalVerificationRequest) (*models.SignalVerification, *responseError) {
	if err := utils.NewValidator().Struct(svr); err != nil {
		return nil, &responseError{
			Description: &fiber.Map{
				"msg": utils.ValidationErrorsToText(err),
			}, StatusCode: fiber.StatusBadRequest}
	}

	true_sv, _ := db.GetSignalVerification(svr.ID)

	if *true_sv == (models.SignalVerification{}) {
		return nil, &responseError{
			Description: &fiber.Map{
				"msg": "Code/ID ist ungültig.",
			}, StatusCode: fiber.StatusBadRequest}
	}

	if true_sv.ValidUntil.Before(time.Now()) {
		if err := db.DeleteSignalVerification(true_sv); err != nil {
			log.Println(err.Error())

			return nil, &responseError{
				Description: &fiber.Map{
					"msg": "Etwas ist schiefgelaufen.",
				}, StatusCode: fiber.StatusBadRequest}
		}

		return nil, &responseError{
			Description: &fiber.Map{
				"msg": "Zeit ist abgelaufen.",
			}, StatusCode: fiber.StatusBadRequest}
	}

	if svr.VerificationCode != true_sv.VerificationCode {
		return nil, &responseError{
			Description: &fiber.Map{
				"msg": "Code/ID ist ungültig.",
			}, StatusCode: fiber.StatusBadRequest}
	}

	/* if err := db.DeleteSignalVerification(true_sv); err != nil {
		log.Println(err.Error())

		return nil, &responseError{
			Description: &fiber.Map{
				"msg": "Something went wrong",
			}, StatusCode: fiber.StatusInternalServerError}
	} */

	return true_sv, nil
}

func CreateAssignment(c *fiber.Ctx) error {
	svr := &SignalVerificationRequest{}
	if err := c.BodyParser(svr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		log.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Etwas ist schiefgelaufen.",
		})
	}

	true_sv, respErr := verifiyRequest(db, c, svr)
	if respErr != nil {
		return c.Status(respErr.StatusCode).JSON(respErr.Description)
	}

	if _, err := db.GetAssignmentByMoodleToken(true_sv.MoodleToken); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Bereits angemeldet. Es passiert nichts.",
		})
	}

	a := models.Assignment{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		MoodleToken: strings.ToLower(true_sv.MoodleToken),
		PhoneNumber: true_sv.PhoneNumber,
		Assignments: models.Assignments{},
	}

	if err := db.CreateAssignment(&a); err != nil {
		log.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Etwas ist schiefgelaufen.",
		})
	}

	return c.JSON(&fiber.Map{
		"msg": "Registrierung erfolgreich!",
	})
}

func DeleteAssignment(c *fiber.Ctx) error {
	svr := &SignalVerificationRequest{}
	if err := c.BodyParser(svr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		log.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Something went wrong",
		})
	}

	true_sv, respErr := verifiyRequest(db, c, svr)
	if respErr != nil {
		return c.Status(respErr.StatusCode).JSON(respErr.Description)
	}

	old_a, err := db.GetAssignmentByMoodleToken(true_sv.MoodleToken)

	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "no entry found",
		})
	}

	if err := db.DeleteAssignment(old_a); err != nil {
		log.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Something went wrong",
		})
	}

	return c.JSON(&fiber.Map{
		"msg": "removed entry",
	})
}
