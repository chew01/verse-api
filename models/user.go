package models

import (
	"context"
	"github.com/chew01/verse-api/db"
	"github.com/chew01/verse-api/utils"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"time"
)

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type PublicUser struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserSignupForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) GetAll() ([]PublicUser, error) {
	pool := db.Pool()
	var users []PublicUser
	rows, err := pool.Query(context.Background(), "SELECT id, username, email, created_at FROM userdata")
	if err != nil {
		return nil, utils.DatabaseQueryErr
	}
	defer rows.Close()

	for rows.Next() {
		var user PublicUser
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, utils.DataScanErr
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

func (u User) GetOneByName(name string) (PublicUser, error) {
	pool := db.Pool()
	var user PublicUser
	row := pool.QueryRow(context.Background(), "SELECT name, username, email, created_at FROM userdata WHERE name = $1", name)
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return user, pgx.ErrNoRows
		}
		return user, utils.DataScanErr
	}
	return user, nil
}

func (u User) Create(form UserSignupForm) error {
	id, err := uuid.NewV4()
	if err != nil {
		return utils.UUIDGenErr
	}

	hash, err := utils.HashPassword(form.Password)
	if err != nil {
		return utils.HashGenErr
	}

	pool := db.Pool()
	if _, err := pool.Exec(context.Background(), "INSERT INTO userdata (id, username, email, password, created_at) VALUES ($1, $2, $3, $4, $5)", id, form.Username, form.Email, hash, time.Now()); err != nil {
		return utils.DatabaseExecErr
	}

	return nil
}

func (u User) GetHashByEmail(email string) (string, error) {
	var hash string
	pool := db.Pool()
	row := pool.QueryRow(context.Background(), "SELECT password FROM userdata WHERE email = $1", email)
	if err := row.Scan(&hash); err != nil {
		return "", utils.DataScanErr
	}
	return hash, nil
}
