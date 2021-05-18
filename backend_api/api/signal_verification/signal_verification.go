package signal_verification

/**
import (
	"math/rand"
	"time"

	"github.com/Dattito/HMN_backend_api/database"
	"github.com/Dattito/HMN_backend_api/methods"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SignalVerificationRequest struct {
	PhoneNumber string `validate:"required,e164"`
	Moodletoken string `validate:"required;len=32"`
}

type SignalVerificationResponse struct {
	Id string
}

type SignalVerification struct {
	gorm.Model
	PhoneNumber string    `gorm:"UNIQUE;NOT NULL" json:"phoneNumber"`
	CodeNumber  int       `gorm:"NOT NULL" json:"codeNumber"`
	ValidUntil  time.Time `gorm:"NOT NULL" json:"validUntil"`
}

var Validate *validator.Validate

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := Validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func PostSignalVerificationV1(c *fiber.Ctx) error {
	svr := new(SignalVerificationRequest)

	if err := c.BodyParser(svr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := ValidateStruct(*svr)
	if errors != nil {
		return c.JSON(errors)
	}

	db := database.DBConn

	signalVerification := SignalVerification{
		PhoneNumber: methods.ToInternationalNumber(svr.PhoneNumber),
		CodeNumber:  1111 + rand.Intn(8888),
		ValidUntil:  time.Now().Add(time.Minute * 5),
	}

	if err := db.Create(&signalVerification); err != nil {
		return c.Status(409).JSON(fiber.Map{"error": "Can't store into database (maybe phoneNumber already exists?)"})
	}

	return c.JSON(signalVerification)
} */
