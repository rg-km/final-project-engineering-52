package repository

import (
	"context"
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
func (db *categoryConn) Fetch(ctx context.Context) ([]models.Category, error) {
	query := `SELECT * FROM categories`

	rows, err := db.conn.QueryContext(ctx, query)
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
func (db *categoryConn) FetchById(ctx context.Context, id int64) (models.Category, error) {
	query := `SELECT * FROM categories WHERE id = ?`

	row := db.conn.QueryRowContext(ctx, query, id)

	var c models.Category

	if err := row.Scan(&c.ID, &c.CategoryName, &c.CreatedAt); err != nil {
		return c, err
	}

	return c, nil
}

// create category
func (ca *categoryConn) Create(ctx context.Context, c *models.Category) (models.Category, error) {
	query := `INSERT INTO categories (category_name) VALUES(?)`

	row, err := ca.conn.ExecContext(ctx, query, &c.CategoryName)
	if err != nil {
		return *c, err
	}

	lastId, _ := row.LastInsertId()

	res, err := ca.FetchById(ctx, lastId)
	if err != nil {
		return *c, err
	}

	return res, nil
}

// TODO: Update

// TODO: Delete
