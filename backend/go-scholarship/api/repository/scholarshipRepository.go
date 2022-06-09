package repository

import (
	"database/sql"

	"go-scholarship/api/models"
)

type scholarConn struct {
	conn *sql.DB
}

func NewScholarshipRepository(conn *sql.DB) models.ScholarshipRepository {
	return &scholarConn{conn}
}

// Fetch
func (db *scholarConn) Fetch() ([]models.Scholarship, error) {
	query := `SELECT * FROM scholarships`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}

	var s []models.Scholarship

	for rows.Next() {
		var scholarship models.Scholarship
		err := rows.Scan(&scholarship.ID, &scholarship.UserID, &scholarship.Name, &scholarship.Description, &scholarship.Image, &scholarship.CategoryID, &scholarship.CreatedAt)

		if err != nil {
			return nil, err
		}

		s = append(s, scholarship)
	}

	return s, nil
}

// TODO: FetchById

// TODO: Create

// TODO: Update

// TODO: Delete
