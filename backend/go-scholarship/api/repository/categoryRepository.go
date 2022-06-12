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

// fetch categories
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

// fetchById category
func (db *categoryConn) FetchById(id int64) (models.Category, error) {
	query := `SELECT * FROM categories WHERE id = ?`

	row := db.conn.QueryRow(query, id)

	var c models.Category

	if err := row.Scan(&c.ID, &c.CategoryName, &c.CreatedAt); err != nil {
		return c, err
	}

	return c, nil
}

// TODO: Create

// TODO: Update

// TODO: Delete

