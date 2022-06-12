package repository

import (
	"context"
	"database/sql"

	"go-scholarship/api/models"
	"go-scholarship/utils/hash"
)

type userConn struct {
	conn *sql.DB
}

func NewUserRepo(conn *sql.DB) models.UserRepository {
	return &userConn{conn}
}

// fetch user by email
func (repo *userConn) fetchUserByEmail(ctx context.Context, email string) (models.User, error) {
	var u models.User
	sqlStmt := `SELECT * FROM users WHERE email = ?`
	row := repo.conn.QueryRow(sqlStmt, email)
	err := row.Scan(&u.ID, &u.Name, &u.Image, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return u, err
	}

	return u, nil
}

// login
func (repo *userConn) Login(ctx context.Context, l *models.Login) (models.User, error) {
	u, err := repo.fetchUserByEmail(ctx, l.Email)
	if err != nil {
		return u, err
	}

	// check if password matches
	if err := hash.CheckPassword(u.Password, l.Password); err != nil {
		return u, err
	}

	return u, nil
}

// register
func (repo *userConn) Register(ctx context.Context, u *models.User) (models.User, error) {
	// check if email already exists
	if _, err := repo.fetchUserByEmail(ctx, u.Email); err == nil {
		return models.User{}, err
	}

	// hash password
	u.Password, _ = hash.HashPassword(u.Password)

	res, err := repo.Create(ctx, u)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

// TODO: logout

// fetch users
func (repo *userConn) Fetch(ctx context.Context) ([]models.User, error) {
	var us []models.User
	query := `SELECT * FROM users`
	rows, err := repo.conn.QueryContext(ctx, query)
	if err != nil {
		return us, err
	}

	defer rows.Close()

	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Name, &u.Image, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
		if err != nil {
			return us, err
		}
		us = append(us, u)
	}

	return us, nil
}

// fetch user by id
func (repo *userConn) FetchById(ctx context.Context, id int64) (models.User, error) {
	var u models.User
	sqlStmt := `SELECT * FROM users WHERE id = ?`
	row := repo.conn.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(&u.ID, &u.Name, &u.Image, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return u, err
	}

	return u, nil
}

// create user
func (repo *userConn) Create(ctx context.Context, u *models.User) (models.User, error) {
	// hash password
	u.Password, _ = hash.HashPassword(u.Password)

	// check if email already exists
	if _, err := repo.fetchUserByEmail(ctx, u.Email); err == nil {
		return *u, err
	}

	query := `INSERT INTO users (name, image, email, password) VALUES(?, ?, ?, ?)`

	row, err := repo.conn.ExecContext(ctx, query, &u.Name, &u.Image, &u.Email, &u.Password)
	if err != nil {
		return *u, err
	}

	lastId, _ := row.LastInsertId()

	res, err := repo.FetchById(ctx, lastId)
	if err != nil {
		return *u, err
	}

	return res, nil
}

// update user
func (repo *userConn) Update(ctx context.Context, id int64, u *models.User) (models.User, error) {
	// check the user if exists
	res, err := repo.FetchById(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	// hash password
	password, _ := hash.HashPassword(u.Password)

	query := `UPDATE users SET name = ?, image = ?,  email = ?, password = ? WHERE id = ?`

	_, err = repo.conn.ExecContext(ctx, query, &u.Name, &u.Image, &u.Email, &password, id)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

// delete user
func (repo *userConn) Delete(ctx context.Context, id int64) error {
	// check the user if exists
	if _, err := repo.FetchById(ctx, id); err != nil {
		return err
	}

	query := `DELETE FROM users WHERE id = ?`
	_, err := repo.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
