package models

import (
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID int64
	Name string `binding:"required"`
	Email string `binding:"required"`
	Password string `binding:"required"`
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