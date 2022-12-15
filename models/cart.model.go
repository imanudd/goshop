package models

import (
	"Go-Shop/connection"
	"Go-Shop/helper"
	"fmt"
	"net/http"
	"time"
)

type Cart struct {
	Id          int       `json:"id"`
	User_Id     int       `json:"user_id"`
	Product_id  int       `json:"product_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Category_ID int       `json:"category"`
	ProductName string    `json:"product_name"`
	Weight      string    `json:"weight"`
	Price       int       `json:"price"`
	Qty         int       `json:"qty"`
	Total       int       `json:"total"`
	CreatedAt   time.Time `json:"created_at"`
}

var con = connection.Init()
var res helper.Response

func AddToCart(User_id, Qty, Product_id int) (helper.Response, error) {
	var price int
	var cart Cart
	totalquery := "SELECT price FROM products where id = $1"

	err := con.QueryRow(totalquery, Product_id).Scan(&price)
	cart.Total = Qty * price
	fmt.Println("USER ID  ", User_id)
	sqlStatment := "INSERT INTO cart (user_id,product_id,qty,created_at,total) VALUES ($1,$2,$3,$4,$5)"
	stmt, err := con.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&User_id, &Product_id, &Qty, time.Now(), &cart.Total)

	if err != nil {
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "sukses"
	return res, nil
}

func ShowMyCart(id int) (helper.Response, error) {
	var cart Cart
	var arrcart []Cart

	fieldselected := "users.username,users.email,products.category_id,products.product_name,products.price,products.weight,cart.id,cart.user_id,cart.qty"
	sqlStatement := "SELECT " + fieldselected + "  FROM cart JOIN users ON cart.user_id = users.id JOIN products ON cart.product_id = products.id where user_id = $1 ORDER BY id desc"

	rows, err := con.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&cart.Username, &cart.Email, &cart.Category_ID, &cart.ProductName, &cart.Price, &cart.Weight, &cart.Id, &cart.User_Id, &cart.Qty)
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

func DeleteCart(id, user_id int) (helper.Response, error) {

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
