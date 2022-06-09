package repository

import (
	"database/sql"

	"go-scholarship/api/models"
)

type commentConn struct {
	conn *sql.DB
}

func NewCommentRepository(db *sql.DB) models.CommentRepository {
	return &commentConn{db}
}

// TODO: FetchComments
func (db *commentConn) Fetch() ([]models.Comment, error) {
	query := `SELECT * FROM comments`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cs []models.Comment

	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.ID, &c.Content, &c.UserID, &c.ScholarshipID, &c.CreatedAt); err != nil {
			return nil, err
		}

		cs = append(cs, c)
	}

	return cs, nil
}

// TODO: FetchCommentById

// TODO: Create

// TODO: Update

// TODO: Delete
