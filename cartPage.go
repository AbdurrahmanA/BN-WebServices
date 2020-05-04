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
			writeResponse(w, succesfullyRecordedError())
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func cartPageControl(orders string) {
	//{"order_status":"0","orders":[{"product_id":"5e4ef7b315a9c347b8810e66","product_name":"beaconCUK","product_price":"0.95"}],"payment_type":"CARDXXX","total_price":"3.1","contact_info":{"user_real_name":"xxxxxxxxxxxx","user_id":"5e4d83a99510103a245aec83","user_surname":"xxxxxxxxxxxxxxxxx","user_address":"xxxx 07 xxxx a CR007","user_phone":"+90533345646546"}}

}
