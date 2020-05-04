package main

import (
	"net/http"
)

func cartPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("orders") == "" {
			writeResponse(w, requiredInputError("orders "))
		} else if r.FormValue("user_real_name") == "" {
			writeResponse(w, requiredInputError("Kullanıcı Adı"))
		} else if r.FormValue("user_id") == "" {
			writeResponse(w, requiredInputError("Kullanıcı ID "))
		} else if r.FormValue("user_surname") == "" {
			writeResponse(w, requiredInputError("Kullanıcı Soyadı "))
		} else if r.FormValue("user_address") == "" {
			writeResponse(w, requiredInputError("Adres "))
		} else if r.FormValue("user_phone") == "" {
			writeResponse(w, requiredInputError("Telefon "))
		} else {
			succesfullyRecordedError()
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func cartPageControl(orders string) {

}
