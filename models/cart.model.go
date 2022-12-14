package models

import (
	"Go-Shop/connection"
	"net/http"
	"time"
)

type Cart struct {
	Id          int       `json:"id"`
	User_id     int       `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Category    string    `json:"category"`
	ProductName string    `json:"productname"`
	Weight      string    `json:"weight"`
	Price       int       `json:"price"`
	Qty         int       `json:"qty"`
	Total       int       `json:"total"`
	CreatedAt   time.Time `json:"createdat"`
}

func AddToCart(id, qty, product_id int) (Response, error) {
	var res Response
	var price int
	con := connection.ConnectDB()
	totalquery := "SELECT price FROM products where id = $1"

	err := con.QueryRow(totalquery, product_id).Scan(&price)
	total := qty * price

	sqlStatment := "INSERT INTO cart (user_id,product_id,qty,created_at,total) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	stmt, err := con.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(&id, &product_id, &qty, time.Now(), &total).Scan(&id)

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "sukses"

	return res, nil
}

func ShowMyCart(id int) (Response, error) {
	var cart Cart
	var arrcart []Cart
	var res Response
	con := connection.ConnectDB()
	defer con.Close()

	fieldselected := "users.username,users.email,products.category,products.product_name,products.price,products.weight,cart.id,cart.user_id,cart.qty"
	sqlStatement := "SELECT " + fieldselected + "  FROM cart JOIN users ON cart.user_id = users.id JOIN products ON cart.product_id = products.id where user_id = $1 ORDER BY id desc"

	rows, err := con.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&cart.Username, &cart.Email, &cart.Category, &cart.ProductName, &cart.Price, &cart.Weight, &cart.Id, &cart.User_id, &cart.Qty)
		if err != nil {
			return res, err
		}
		cart.Total = cart.Price * cart.Qty
		arrcart = append(arrcart, cart)

	}
	if arrcart == nil {
		res.Message = "Product Kosong atau kesalahan input category"
		res.Data = ""
		return res, err
	} else {
		res.Status = http.StatusOK
		res.Message = "Succsess"
		res.Data = arrcart
	}
	return res, nil
}

func DeleteCart(id, user_id int) (Response, error) {
	var res Response

	con := connection.ConnectDB()
	defer con.Close()

	sqlStatement := "DELETE FROM cart WHERE id = $1 AND user_id = $2"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, user_id)
	if err != nil {
		return res, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES DELETE CART"
	res.Data = map[string]int64{
		"rowaffected": rowAffected,
	}

	return res, nil

}
