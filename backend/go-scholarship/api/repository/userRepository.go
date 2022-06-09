package repository

import (
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
func (repo *userConn) fetchUserByEmail(email string) (models.User, error) {
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
func (repo *userConn) Login(l *models.Login) (models.User, error) {
	u, err := repo.fetchUserByEmail(l.Email)
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
func (repo *userConn) Register(u *models.User) (models.User, error) {
	// check if email already exists
	if _, err := repo.fetchUserByEmail(u.Email); err == nil {
		return models.User{}, err
	}

	// hash password
	u.Password, _ = hash.HashPassword(u.Password)

	res, err := repo.Create(u)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

// TODO: logout

// fetch users
func (repo *userConn) Fetch() ([]models.User, error) {
	var us []models.User
	query := `SELECT * FROM users`
	rows, err := repo.conn.Query(query)
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
func (repo *userConn) FetchById(id int64) (models.User, error) {
	var u models.User
	sqlStmt := `SELECT * FROM users WHERE id = ?`
	row := repo.conn.QueryRow(sqlStmt, id)
	err := row.Scan(&u.ID, &u.Name, &u.Image, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return u, err
	}

	return u, nil
}

// create user
func (repo *userConn) Create(u *models.User) (models.User, error) {
	// hash password
	u.Password, _ = hash.HashPassword(u.Password)

	// check if email already exists
	if _, err := repo.fetchUserByEmail(u.Email); err == nil {
		return *u, err
	}

	query := `INSERT INTO users (name, image, email, password) VALUES(?, ?, ?, ?)`

	row, err := repo.conn.Exec(query, &u.Name, &u.Image, &u.Email, &u.Password)
	if err != nil {
		return *u, err
	}

	lastId, _ := row.LastInsertId()

	res, err := repo.FetchById(lastId)
	if err != nil {
		return *u, err
	}

	return res, nil
}

// update user
func (repo *userConn) Update(id int64, u *models.User) (models.User, error) {
	// check the user if exists
	res, err := repo.FetchById(id)
	if err != nil {
		return models.User{}, err
	}

	// hash password
	password, _ := hash.HashPassword(u.Password)

	query := `UPDATE users SET name = ?, image = ?,  email = ?, password = ? WHERE id = ?`

	_, err = repo.conn.Exec(query, &u.Name, &u.Image, &u.Email, &password, id)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

// delete user
func (repo *userConn) Delete(id int64) error {
	// check the user if exists
	if _, err := repo.FetchById(id); err != nil {
		return err
	}

	query := `DELETE FROM users WHERE id = ?`
	_, err := repo.conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
