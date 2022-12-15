package models

import (
	"Go-Shop/connection"
	"Go-Shop/helper"
	"net/http"
	"time"
)

type Transaction struct {
	Id         int       `json:"id"`
	Cart_id    int       `json:"cart_id"`
	Payment_id int       `json:"payment_id"`
	Created_at time.Time `json:"created_at"`
	Total      int       `json:"total"`
}

func AddTransaction(user_id, payment_id int) (helper.Response, error) {
	var total int
	var con = connection.Init()
	var res helper.Response
	totalquery := "SELECT sum(total) as total from cart where user_id = $1"
	err = con.QueryRow(totalquery, user_id).Scan(&total)

	sqlStatment := "INSERT INTO transactions (user_id,payment_id,created_at,total) VALUES ($1,$2,$3,$4)"
	stmt, err := con.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id, payment_id, time.Now(), total)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "sukses"
	return res, nil

}
