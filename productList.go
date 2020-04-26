package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func productsListPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var products, control = productsList(r.FormValue("token"))
		if products == nil {
			if control == "Nil" {
				writeResponse(w, notFindRecordError())
			} else {
				writeResponse(w, someThingWentWrong())
			}
		} else {
			writeResponse(w, string(products))
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}
func productsList(token string) ([]byte, string) {
	var data []byte
	var l []*ProductsInApp
	var prs *ProductsInApp
	product := &Products{}
	products := connection.Collection("products").Find(bson.D{})
	for products.Next(product) {
		prs = &ProductsInApp{product.Id, product.ProductDescription, product.ProductName, product.ProductPrice, product.ProductType}
		l = append(l, prs)
	}
	data, _ = json.Marshal(l)
	if l == nil {
		return nil, "Nil"
	}
	if l != nil {
		response := &ProductsArray{l}
		data, _ = json.Marshal(response)
		return addError(data), ""
	}
	return data, ""
}
