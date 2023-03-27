package models

import (
	"Go-Shop/connection"
	"Go-Shop/helper"
	"net/http"
	"time"
)

type Transaction struct {
	Id         int       `json:"id"`
	User_id    int       `json:"cart_id"`
	Payment_id int       `json:"payment_id"`
	Created_at time.Time `json:"created_at"`
	Total      int       `json:"total"`
}

func AddTransaction(user_id, payment_id int) (helper.Response, error) {
	var total int
	var transaction Transaction
	var con = connection.Init()
	var res helper.Response
	totalquery := "SELECT sum(total) as total from cart where user_id = $1"
	err = con.QueryRow(totalquery, user_id).Scan(&total)

	sqlStatment := "INSERT INTO transactions (user_id,payment_id,created_at,total) VALUES ($1,$2,$3,$4) RETURNING id"
	stmt, err := con.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(user_id, payment_id, time.Now(), total).Scan(&transaction.Id)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "sukses"
	res.Data = map[string]int{
		"Id Transaction": transaction.Id,
		"Payment Id":     payment_id,
	}
	return res, nil

}
