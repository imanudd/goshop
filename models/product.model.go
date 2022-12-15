package models

import (
	"Go-Shop/connection"
	"Go-Shop/helper"
	"net/http"
)

type Product struct {
	Id          int    `json:"id"`
	Category_Id string `json:"category_id"`
	ProductName string `json:"product_name"`
	Weight      string `json:"weight"`
	Price       int    `json:"price"`
}

var err error

func SearchByCategory(category_id int) (helper.Response, error) {
	var product Product
	var arrproduct []Product
	var con = connection.Init()
	var res helper.Response

	sqlStatement := "SELECT * FROM products WHERE category_id = $1"

	rows, err := con.Query(sqlStatement, category_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Category_Id, &product.ProductName, &product.Weight, &product.Price)
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
