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

// fetch user by id for comparing password
func (u *userConn) fetchById(ctx context.Context, id int64) (models.User, error) {
	var user models.User
	sqlStmt := `SELECT * FROM users WHERE id = ?`
	row := u.conn.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(&user.ID, &user.Name, &user.Image, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// login
func (u *userConn) Login(ctx context.Context, login *models.Login) (models.UserResponse, error) {
	user, err := u.fetchUserByEmail(ctx, login.Email)
	if err != nil {
		return models.UserResponse{}, err
	}

	// check if password matches
	if err := hash.CheckPassword(user.Password, login.Password); err != nil {
		return models.UserResponse{}, err
	}

	userResponse := &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Image:     user.Image,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	return *userResponse, nil
}

// register
func (u *userConn) Register(ctx context.Context, user *models.User) (models.UserResponse, error) {
	res, err := u.Create(ctx, user)
	if err != nil {
		return models.UserResponse{}, err
	}

	return res, nil
}

// fetch users
func (u *userConn) Fetch(ctx context.Context) ([]models.UserResponse, error) {
	query := `SELECT * FROM users`
	rows, err := u.conn.QueryContext(ctx, query)
	if err != nil {
		return []models.UserResponse{}, err
	}

	defer rows.Close()

	var users []models.UserResponse
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Image, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
		if err != nil {
			return []models.UserResponse{}, err
		}

		userResponse := &models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Image:     user.Image,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		}

		users = append(users, *userResponse)
	}

	return users, nil
}

// fetch user by id
func (u *userConn) FetchById(ctx context.Context, id int64) (models.UserResponse, error) {
	var user models.User
	sqlStmt := `SELECT * FROM users WHERE id = ?`
	row := u.conn.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(&user.ID, &user.Name, &user.Image, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return models.UserResponse{}, err
	}

	userResponse := &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Image:     user.Image,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	return *userResponse, nil
}

// create user
func (u *userConn) Create(ctx context.Context, user *models.User) (models.UserResponse, error) {
	// hash password
	user.Password, _ = hash.HashPassword(user.Password)

	query := `INSERT INTO users (name, image, email, password) VALUES(?, ?, ?, ?)`

	row, err := u.conn.ExecContext(ctx, query, &user.Name, &user.Image, &user.Email, &user.Password)
	if err != nil {
		return models.UserResponse{}, err
	}

	lastId, _ := row.LastInsertId()

	res, err := u.FetchById(ctx, lastId)
	if err != nil {
		return models.UserResponse{}, err
	}

	userResponse := &models.UserResponse{
		ID:        res.ID,
		Name:      res.Name,
		Image:     res.Image,
		Email:     res.Email,
		Role:      res.Role,
		CreatedAt: res.CreatedAt,
	}

	return *userResponse, nil
}

// update user
func (u *userConn) Update(ctx context.Context, id int64, user *models.User) (models.UserResponse, error) {
	usr, err := u.fetchById(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	// compare with the old password
	if user.Password != usr.Password {
		// hash password
		user.Password, _ = hash.HashPassword(user.Password)
	}

	query := `UPDATE users SET name = ?, image = ?,  email = ?, password = ? WHERE id = ?`

	_, err = u.conn.ExecContext(ctx, query, &user.Name, &user.Image, &user.Email, &user.Password, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	res, err := u.FetchById(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return res, nil
}

// delete user
func (u *userConn) Delete(ctx context.Context, id int64) error {
	// check the user if exists
	if _, err := u.FetchById(ctx, id); err != nil {
		return err
	}

	query := `DELETE FROM users WHERE id = ?`
	_, err := u.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
