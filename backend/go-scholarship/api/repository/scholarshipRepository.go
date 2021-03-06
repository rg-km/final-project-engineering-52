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
func (s *scholarConn) Fetch(ctx context.Context) ([]models.ScholarResponse, error) {
	query := `SELECT * FROM scholarships`

	rows, err := s.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var scholars []models.ScholarResponse

	for rows.Next() {
		var scholarship models.ScholarResponse
		err := rows.Scan(&scholarship.ID, &scholarship.Name, &scholarship.Description, &scholarship.Image, &scholarship.Category.ID, &scholarship.User.ID, &scholarship.CreatedAt)

		if err != nil {
			return nil, err
		}

		scholars = append(scholars, scholarship)
	}

	return scholars, nil
}

// fetch by id scholarship
func (s *scholarConn) FetchById(ctx context.Context, id int64) (models.ScholarResponse, error) {
	query := `SELECT * FROM scholarships WHERE id = ?`

	row := s.conn.QueryRowContext(ctx, query, id)

	var scholar models.ScholarResponse
	if err := row.Scan(&scholar.ID, &scholar.Name, &scholar.Description, &scholar.Image, &scholar.Category.ID, &scholar.User.ID, &scholar.CreatedAt); err != nil {
		return models.ScholarResponse{}, err
	}

	return scholar, nil
}

// create scholarship
func (s *scholarConn) Create(ctx context.Context, scholarReq *models.ScholarRequest) (models.ScholarResponse, error) {
	query := `INSERT INTO scholarships (name, description, image, category_id, user_id) VALUES(?, ?, ?, ?, ?)`

	row, err := s.conn.ExecContext(ctx, query, &scholarReq.Name, &scholarReq.Description, &scholarReq.Image, &scholarReq.CategoryID, &scholarReq.UserID)
	if err != nil {
		return models.ScholarResponse{}, err
	}

	lastId, _ := row.LastInsertId()

	res, err := s.FetchById(ctx, lastId)
	if err != nil {
		return models.ScholarResponse{}, err
	}

	return res, nil
}

// update scholarship
func (s *scholarConn) Update(ctx context.Context, id int64, scholarReq *models.ScholarRequest) (models.ScholarResponse, error) {
	_, err := s.FetchById(ctx, scholarReq.ID)
	if err != nil {
		return models.ScholarResponse{}, err
	}

	query := `UPDATE scholarships SET name = ?, description = ?, image = ?, category_id = ?, user_id = ? WHERE id = ?`

	_, err = s.conn.ExecContext(ctx, query, &scholarReq.Name, &scholarReq.Description, &scholarReq.Image, &scholarReq.CategoryID, &scholarReq.UserID, id)
	if err != nil {
		return models.ScholarResponse{}, err
	}

	res, err := s.FetchById(ctx, id)
	if err != nil {
		return models.ScholarResponse{}, err
	}

	return res, nil
}

// delete scholarship
func (s *scholarConn) Delete(ctx context.Context, id int64) error {
	// check scholar if exist
	_, err := s.FetchById(ctx, id)
	if err != nil {
		return err
	}

	query := `DELETE FROM scholarships WHERE id = ?`

	_, err = s.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
