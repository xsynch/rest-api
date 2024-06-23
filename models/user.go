package models

import (
	"errors"

	"udemy.com/rest-api/db"
	"udemy.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	password string `binding:"required"`
}

func (u User) Save() error {
	query := "insert into users(email, password) values (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err 
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err 
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err 

}

func (u *User) ValidateCredentials() error {
	query := "Select id, password from users where email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID,&retrievedPassword)
	if err != nil {
		return errors.New("Password is invalid")
	}
	validPass := utils.CheckPasswordHash(u.password, retrievedPassword)
	if !validPass {
		return errors.New("Password is invalid")
	}
	return nil 
}