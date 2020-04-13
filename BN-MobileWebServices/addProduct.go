package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func addProductPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
		} else {
			var devices, control = addProduct(r.FormValue("token"), r.FormValue("proDes"), r.FormValue("proName"), r.FormValue("proPrice"), r.FormValue("proType"))
			if devices == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "ÜrünFiyatı" {
					writeResponse(w, incorrectInput("Ürün Fiyatı"))
				} else if control == "Ürüntipi" {
					writeResponse(w, incorrectInput("Ürün tipi"))
				} else if control == "Token" {
					writeResponse(w, invalidPermission())
				} else if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func addProduct(token string, proDes string, proName string, proPrice string, proType string) (bool, string) {
	proPrices, err := strconv.ParseFloat(proPrice, 32)
	proPriceFloat := float32(proPrices)
	if err != nil {
		return false, "ÜrünFiyatı"
	}
	proTypeInt, err := strconv.Atoi(proType)
	if err != nil {
		return false, "Ürüntipi"
	}
	perForAdd := checkPermission(token)
	if perForAdd == false {
		return false, "Token"
	}
	device := &ProductsForAdd{
		ProductDescription: proDes,
		ProductName:        proName,
		ProductPrice:       proPriceFloat,
		ProductType:        proTypeInt,
	}
	errs := connection.Collection("products").Save(device)
	if errs != nil {
		fmt.Println(errs.Error())
		return false, "Save"
	}
	return true, ""
}
