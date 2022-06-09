package repository

import (
	"database/sql"

	"go-scholarship/api/models"
)

type categoryConn struct {
	conn *sql.DB
}

func NewCategoryRepository(db *sql.DB) models.CategoryRepository {
	return &categoryConn{db}
}

func (db *categoryConn) Fetch() ([]models.Category, error) {
	query := `SELECT * FROM categories`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cs []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.CategoryName, &c.CreatedAt); err != nil {
			return nil, err
		}

		cs = append(cs, c)
	}

	return cs, nil
}

// TODO: FetchById

// TODO: Create

// TODO: Update

// TODO: Delete

