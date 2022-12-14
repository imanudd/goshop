package models

import (
	"Go-Shop/connection"
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

func AddTransaction(user_id, payment_id int) (Response, error) {
	var res Response
	var total int
	con := connection.ConnectDB()
	totalquery := "SELECT sum(total) as total from cart where user_id = $1"
	err = con.QueryRow(totalquery, user_id).Scan(&total)

	sqlStatment := "INSERT INTO transactions (user_id,payment_id,created_at,total) VALUES ($1,$2,$3,$4)"
	stmt, err := con.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	// err = stmt.QueryRow(&cart_id, &payment_id, time.Now(), &total)
	_, err = stmt.Exec(user_id, payment_id, time.Now(), total)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "sukses"
	return res, nil

}
