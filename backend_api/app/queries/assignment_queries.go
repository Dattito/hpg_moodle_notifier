package queries

import (
	"github.com/Dattito/HMN_backend_api/app/models"
	"github.com/jmoiron/sqlx"
)

type AssignmentQueries struct {
	*sqlx.DB
}

func (s *AssignmentQueries) GetAssignmentByMoodleToken(moodle_token string) (*models.Assignment, error) {
	a := models.Assignment{}

	query := `SELECT * FROM assignments WHERE moodle_token = $1`

	if err := s.Get(&a, query, moodle_token); err != nil {
		return &a, err
	}

	return &a, nil
}

func (s *AssignmentQueries) CreateAssignment(a *models.Assignment) error {
	query := `INSERT INTO assignments (id, created_at, updated_at, moodle_token, phone_number, assignments) VALUES ($1, $2, $3, $4, $5, $6)`

	if _, err := s.Exec(query, a.ID, a.CreatedAt, a.UpdatedAt, a.MoodleToken, a.PhoneNumber, a.Assignments); err != nil {
		return err
	}
	return nil
}

func (s *AssignmentQueries) DeleteAssignment(a *models.Assignment) error {
	query := `DELETE FROM assignments WHERE id = $1`

	if _, err := s.Exec(query, a.ID); err != nil {
		return err
	}
	return nil
}
