package models

import (
	"Go-Shop/connection"
	"fmt"
	"net/http"
	"strings"
)

type Product struct {
	Id          int    `json:"id"`
	Category    string `json:"category"`
	ProductName string `json:"productname"`
	Weight      string `json:"weight"`
	Price       int    `json:"price"`
}

var err error

func SearchByCategory(category string) (Response, error) {
	var product Product
	var arrproduct []Product
	var res Response
	con := connection.ConnectDB()
	defer connection.ConnectDB().Close()
	category = strings.ToUpper(category)

	fmt.Println(category)
	sqlStatement := "SELECT * FROM products WHERE category = $1"

	rows, err := con.Query(sqlStatement, category)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Category, &product.ProductName, &product.Weight, &product.Price)
		if err != nil {
			return res, err
		}
		arrproduct = append(arrproduct, product)
	}
	if arrproduct == nil {
		res.Message = "Product Kosong atau kesalahan input category"
		res.Data = ""
		return res, err
	} else {
		res.Status = http.StatusOK
		res.Message = "Succsess"
		res.Data = arrproduct
	}
	return res, nil
}
