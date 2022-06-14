package repository

import (
	"context"
	"database/sql"

	"go-scholarship/api/models"
)

type scholarConn struct {
	conn *sql.DB
}

func NewScholarshipRepository(conn *sql.DB) models.ScholarshipRepository {
	return &scholarConn{conn}
}

// fetch all scholarships
func (s *scholarConn) Fetch(ctx context.Context) ([]models.Scholarship, error) {
	query := `SELECT * FROM scholarships`

	rows, err := s.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var scholars []models.Scholarship

	for rows.Next() {
		var scholarship models.Scholarship
		err := rows.Scan(&scholarship.ID, &scholarship.Name, &scholarship.Description, &scholarship.Image, &scholarship.Category.ID, &scholarship.User.ID, &scholarship.CreatedAt)

		if err != nil {
			return nil, err
		}

		scholars = append(scholars, scholarship)
	}

	return scholars, nil
}

// fetch by id scholarship
func (s *scholarConn) FetchById(ctx context.Context, id int64) (models.Scholarship, error) {
	query := `SELECT * FROM scholarships WHERE id = ?`

	row := s.conn.QueryRowContext(ctx, query, id)

	var scholar models.Scholarship
	if err := row.Scan(&scholar.ID, &scholar.Name, &scholar.Description, &scholar.Image, &scholar.Category.ID, &scholar.User.ID, &scholar.CreatedAt); err != nil {
		return models.Scholarship{}, err
	}

	return scholar, nil
}

// create scholarship
func (s *scholarConn) Create(ctx context.Context, scholar *models.Scholarship) (models.Scholarship, error) {
	query := `INSERT INTO scholarships (name, description, image, category_id, user_id) VALUES(?, ?, ?, ?, ?)`

	row, err := s.conn.ExecContext(ctx, query, &scholar.Name, &scholar.Description, &scholar.Image, &scholar.Category.ID, &scholar.User.ID)
	if err != nil {
		return *scholar, err
	}

	lastId, _ := row.LastInsertId()

	res, err := s.FetchById(ctx, lastId)
	if err != nil {
		return *scholar, err
	}

	return res, nil
}

// update scholarship
func (s *scholarConn) Update(ctx context.Context, id int64, scholar *models.Scholarship) (models.Scholarship, error) {
	query := `UPDATE scholarships SET name = ?, description = ?, image = ?, category_id = ?, user_id = ?`

	_, err := s.conn.ExecContext(ctx, query, &scholar.Name, &scholar.Description, &scholar.Image, &scholar.Category.ID, &scholar.User.ID)
	if err != nil {
		return *scholar, err
	}

	res, err := s.FetchById(ctx, id)
	if err != nil {
		return *scholar, err
	}

	return res, nil
}

// delete scholarship
func (s *scholarConn) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM scholarships WHERE id = ?`

	_, err := s.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
