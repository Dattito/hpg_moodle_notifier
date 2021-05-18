package models

import (
	"time"

	"github.com/google/uuid"
)

type SignalVerification struct {
	ID               uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt        time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt        time.Time `db:"updated_at" json:"updatedAt"`
	MoodleToken      string    `db:"moodle_token" json:"moodleToken" validate:"required,len=32"`
	PhoneNumber      string    `db:"phone_number" json:"phoneNumber" validate:"required,e164"`
	VerificationCode int       `db:"verification_code" json:"verificationCode" validate:"required,min=111111,max=999999"`
	ValidUntil       time.Time `db:"valid_until" json:"validUntil" validate:"required"`
}
