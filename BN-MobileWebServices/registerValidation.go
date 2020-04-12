package main

import (
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func validationRegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.FormValue("token") == "" {
			writeResponse(w, requiredInputError("Anahtar"))
		} else {
			var register, control = validationRegister(r.FormValue("token"))
			if register == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "NotFound" {
					writeResponse(w, notFindRecordError())

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

func validationRegister(token string) (bool, string) {
	person := &Person{}
	err := connection.Collection("users").FindOne(bson.M{"user_info.user_token": token}, person)
	if err != nil {
		return false, "NotFound"
	}
	person.UserInfos.RoleLvl = 1
	errs := connection.Collection("users").Save(person)
	if errs != nil {
		fmt.Println(errs.Error())
		return false, "Save"
	}
	return true, ""

}
