package models

import (
	"api-echo/db"
	"database/sql"
	"fmt"

	helpers "api-echo/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	conn := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := conn.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Username Not found")
		return false, err
	}
	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPassword(password, pwd)
	if !match {
		fmt.Println("Doesn't Match")
		return false, err
	}

	return true, nil
}
