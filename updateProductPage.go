package main

import (
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
)

func updateProductPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" || r.Method == "POST" {
		if r.FormValue("token") == "" {
			writeResponse(w, requiredInputError("Token "))
		} else if r.FormValue("proDes") == "" {
			writeResponse(w, requiredInputError("Ürün açıklaması"))
		} else if r.FormValue("proName") == "" {
			writeResponse(w, requiredInputError("Ürün ismi "))
		} else if r.FormValue("proPrice") == "" {
			writeResponse(w, requiredInputError("Ürün fiyatı "))
		} else if r.FormValue("proType") == "" {
			writeResponse(w, requiredInputError("Ürün tipi "))
		} else if r.FormValue("id") == "" {
			writeResponse(w, requiredInputError("id"))
		} else {
			var control, str = updateProduct(r.FormValue("token"), r.FormValue("proDes"), r.FormValue("proName"), r.FormValue("proPrice"), r.FormValue("proType"), r.FormValue("id"))
			if control == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if str == "Perm" {
					writeResponse(w, invalidPermission())
				} else if str == "ÜrünFiyatı" {
					writeResponse(w, incorrectInput("Ürün Fiyatı"))
				} else if str == "Ürüntipi" {
					writeResponse(w, incorrectInput("Ürün tipi"))
				} else if str == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if str == "NotFound" {
					writeResponse(w, notFindRecordError())
				} else if str == "ID" {
					writeResponse(w, incorrectInput("ID"))
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func updateProduct(token string, proDes string, proName string, proPrice string, proType string, productID string) (bool, string) {
	tokenControl := checkPermission(token)

	if tokenControl == true {
		product := &Products{}
		conroltID, errID := checkObjID(productID)
		if errID == true {
			err := connection.Collection("products").FindById(bson.ObjectIdHex(conroltID), product)
			if err != nil {
				return false, "NotFound"
			}
			proPrices, err := strconv.ParseFloat(proPrice, 32)
			proPriceFloat := float32(proPrices)
			if err != nil {
				return false, "ÜrünFiyatı"
			}
			proTypeInt, err := strconv.Atoi(proType)
			if err != nil {
				return false, "Ürüntipi"
			}
			if proTypeInt > 4 || proTypeInt < 0 {
				return false, "Ürüntipi"
			}
			product.ProductDescription = proDes
			product.ProductName = proName
			product.ProductType = proTypeInt
			product.ProductPrice = proPriceFloat
			errors := connection.Collection("products").Save(product)
			if errors != nil {
				return false, "Save"
			}
			return true, ""
		}
		return false, "ID"
	}
	return false, "Perm"
}
