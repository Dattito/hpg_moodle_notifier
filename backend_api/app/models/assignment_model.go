package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Assignment struct {
	ID          uuid.UUID   `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time   `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time   `db:"updated_at" json:"updatedAt"`
	MoodleToken string      `db:"moodle_token" json:"moodleToken" validate:"required,len=32"`
	PhoneNumber string      `db:"phone_number" json:"phoneNumber" validate:"required,e164"`
	Assignments Assignments `db:"assignments" json:"assignments"`
}

type Assignments []int

func (a *Assignments) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &a)
		return nil
	/*case string:
	json.Unmarshal([]byte(v), &a)
	return nil */
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func (a *Assignments) Value() (driver.Value, error) {
	return json.Marshal(a)
}
