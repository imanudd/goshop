package models

import (
	"Go-Shop/connection"
	"Go-Shop/helper"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
}

func CheckLogin(username, password string) (int, bool, error) {
	var obj User
	var pwd string

	con := connection.ConnectDB()

	sqlStatement := "SELECT id,username,password FROM users WHERE username = $1"

	err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Username, &pwd)

	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found")
		return obj.Id, false, err
	}

	if err != nil {
		fmt.Println("Query ERROR")
		return obj.Id, false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash And Password Doesn't Match !")
		return obj.Id, false, err
	}

	return obj.Id, true, nil
}

func Register(username, password, email, address, phonenumber string) (Response, error) {
	var res Response
	var id int64
	con := connection.ConnectDB()

	sqlStatment := "INSERT INTO users (username,password,email,address,phone_number) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	stmt, err := con.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, password, email, phonenumber, address).Scan(&id)

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "sukses"
	res.Data = map[string]int64{
		"Last_Inserted_id": id,
	}

	return res, nil
}
