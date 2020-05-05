package main

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func passwordChangePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" || r.Method == "POST" {
		if r.FormValue("oldPass") == "" {
			writeResponse(w, requiredInputError("Eski Parola"))
		} else if r.FormValue("newPass") == "" {
			writeResponse(w, requiredInputError("Yeni Parola "))
		} else if r.FormValue("newPassAgain") == "" {
			writeResponse(w, requiredInputError("Yeni Parola Tekrarı"))
		} else if r.FormValue("id") == "" {
			writeResponse(w, requiredInputError("Kullanıcı Numarası"))
		} else {
			var update, control = passwordChange(r.FormValue("oldPass"), r.FormValue("newPass"), r.FormValue("newPassAgain"), r.FormValue("id"))
			if update == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if control == "Nil" {
					writeResponse(w, failedRecordError())
				} else if control == "ID" {
					writeResponse(w, objectIDError())
				} else if control == "OldPass" {
					writeResponse(w, incorrectInput("Eski şifre"))
				} else if control == "OldPassAgain" {
					writeResponse(w, incorrectInput("Eski şifre tekrarı"))
				} else if control == "PassMatch" {
					writeResponse(w, incorrectInput("Yeni şifrelerin uyumu"))
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

func passwordChange(oldPass string, newPass string, newPassAgain string, id string) (bool, string) {
	ID, controlID := checkObjID(id)
	if controlID == true {
		person := &Person{}
		err := connection.Collection("users").FindById(bson.ObjectIdHex(ID), person)
		if err != nil {
			return false, "Nil"
		}
		if person.UserInfos.UserPassword != oldPass {
			return false, "OldPass"
		}
		if newPass != newPassAgain {
			return false, "PassMatch"
		}
		if person.UserInfos.UserPassword == newPass {
			return false, "OldPassAgain"
		}
		person.UserInfos.UserPassword = newPass
		errors := connection.Collection("users").Save(person)
		if errors != nil {
			return false, "Save"
		}
		return true, ""
	}
	return false, "ID"
}
