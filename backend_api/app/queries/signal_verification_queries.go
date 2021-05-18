package queries

import (
	"github.com/Dattito/HMN_backend_api/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SignalVerificationQueries struct {
	*sqlx.DB
}

func (s *SignalVerificationQueries) GetSignalVerification(id uuid.UUID) (*models.SignalVerification, error) {
	sv := models.SignalVerification{}

	query := `SELECT * FROM signal_verifications WHERE id = $1`

	if err := s.Get(&sv, query, id); err != nil {
		return &sv, err
	}

	return &sv, nil
}

func (s *SignalVerificationQueries) GetSignalVerificationByMoodleToken(moodle_token string) (*models.SignalVerification, error) {
	sv := models.SignalVerification{}

	query := `SELECT * FROM signal_verifications WHERE moodle_token = $1`

	if err := s.Get(&sv, query, moodle_token); err != nil {
		return &sv, err
	}

	return &sv, nil
}

func (s *SignalVerificationQueries) CreateSignalVerification(sv *models.SignalVerification) error {
	query := `INSERT INTO signal_verifications VALUES ($1, $2, $3, $4, $5, $6, $7)`

	if _, err := s.Exec(query, sv.ID, sv.CreatedAt, sv.UpdatedAt, sv.MoodleToken, sv.PhoneNumber, sv.VerificationCode, sv.ValidUntil); err != nil {
		return err
	}
	return nil
}

func (s *SignalVerificationQueries) DeleteSignalVerification(sv *models.SignalVerification) error {
	query := `DELETE FROM signal_verifications WHERE id = $1`

	if _, err := s.Exec(query, sv.ID); err != nil {
		return err
	}
	return nil
}
