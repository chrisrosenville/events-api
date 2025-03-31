package models

import (
	"errors"
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID int64
	Name string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users(name, email, password)
	VALUES (?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query);
	if err != nil {
        return err
    }

	defer stmt.Close()

	hash, err := utils.HashPassword(u.Password)
	if err != nil {
        return err
    }

	result, err := stmt.Exec(u.Name, u.Email, hash)
	if err != nil {
        return err
    }

	userId, err := result.LastInsertId()
	u.ID = userId

    return err
}

func (u *User) ValidateCredentials() (error) {
	query := `
	SELECT id, name, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &u.Name, &retrievedPassword)
	if err != nil {
		return err
	}

	isPasswordMatch := utils.CompareHashedPassword(u.Password, retrievedPassword)
    if !isPasswordMatch {
		return errors.New("invalid credentials")
	}

	return nil
}