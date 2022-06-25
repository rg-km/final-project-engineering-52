package repository

import (
	"context"
	"database/sql"

	"go-scholarship/api/models"
)

type commentConn struct {
	conn *sql.DB
}

func NewCommentRepository(db *sql.DB) models.CommentRepository {
	return &commentConn{db}
}

// fetch comments
func (co *commentConn) Fetch(ctx context.Context) ([]models.CommentResponse, error) {
	query := `SELECT * FROM comments`

	rows, err := co.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []models.CommentResponse

	for rows.Next() {
		var comment models.CommentResponse
		if err := rows.Scan(&comment.ID, &comment.Content, &comment.User.ID, &comment.ScholarshipID, &comment.CreatedAt); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

// fetch by id comment
func (co *commentConn) FetchById(ctx context.Context, id int64) (models.CommentResponse, error) {
	query := `SELECT * FROM comments WHERE id = ?`

	row := co.conn.QueryRowContext(ctx, query, id)

	var comment models.CommentResponse
	if err := row.Scan(&comment.ID, &comment.Content, &comment.User.ID, &comment.ScholarshipID, &comment.CreatedAt); err != nil {
		return comment, err
	}

	return comment, nil
}

// create comment
func (co *commentConn) Create(ctx context.Context, commentReq *models.CommentRequest) (models.CommentResponse, error) {
	var commentResp models.CommentResponse
	query := `INSERT INTO comments (content, user_id, scholarship_id) VALUES(?, ?, ?)`

	row, err := co.conn.ExecContext(ctx, query, &commentReq.Content, &commentReq.UserID, &commentReq.ScholarshipID)
	if err != nil {
		return commentResp, err
	}

	userId, _ := row.LastInsertId()

	res, err := co.FetchById(ctx, userId)
	if err != nil {
		return commentResp, err
	}

	return res, nil
}

// update comment
func (co *commentConn) Update(ctx context.Context, id int64, commentReq *models.CommentRequest) (models.CommentResponse, error) {
	var commentResp models.CommentResponse
	query := `UPDATE comments SET content = ?, user_id = ?, scholarship_id = ? WHERE id = ?`

	_, err := co.conn.ExecContext(ctx, query, &commentReq.Content, &commentReq.UserID, &commentReq.ScholarshipID, id)
	if err != nil {
		return commentResp, err
	}

	res, err := co.FetchById(ctx, id)
	if err != nil {
		return commentResp, err
	}

	return res, nil
}

// delete comment
func (co *commentConn) Delete(ctx context.Context, id int64) error {
	// check comment if exist
	if _, err := co.FetchById(ctx, id); err != nil {
		return err
	}

	query := `DELETE FROM comments WHERE id = ?`

	_, err := co.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
