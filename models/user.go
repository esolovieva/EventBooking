package models

import (
	"errors"

	"example.com/eventbooking/models/db"
	"example.com/eventbooking/utils"
)

type User struct {
	ID       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return errors.New("Credentials are invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Credentials are invalid")
	}
	return nil
}
