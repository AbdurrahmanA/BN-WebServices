package main

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func validationLostPasswordPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.FormValue("token") == "" {
			writeResponse(w, requiredInputError("Anahtar"))
		} else if r.FormValue("email") == "" {
			writeResponse(w, requiredInputError("email"))
		} else {
			var lostPass, control = validationLostPassword(r.FormValue("token"), r.FormValue("email"))
			if lostPass == true {
				writeResponse(w, succesfullyError())
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

func validationLostPassword(token string, email string) (bool, string) {
	person := &Person{}
	err := connection.Collection("users").FindOne(bson.M{"user_infos.user_web_token": token, "user_infos.user_mail": email}, person)
	if err != nil {
		return false, "NotFound"
	}
	person.UserInfos.UserWebToken = ""
	errors := connection.Collection("users").Save(person)
	if errors != nil {
		return false, "Save"
	}
	return true, ""
}

// LostPasswordChangePage ss
func LostPasswordChangePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("email") == "" {
			writeResponse(w, requiredInputError("email"))
		} else if r.FormValue("newPass") == "" {
			writeResponse(w, requiredInputError("newPass"))
		} else if r.FormValue("newPassAgain") == "" {
			writeResponse(w, requiredInputError("newPassAgain"))
		} else {
			var lostPass, control = LostPasswordChange(r.FormValue("newPass"), r.FormValue("newPassAgain"), r.FormValue("email"))
			if lostPass == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if control == "OldPassAgain" {
					writeResponse(w, incorrectInput("Eski şifre tekrarı"))
				} else if control == "PassMatch" {
					writeResponse(w, incorrectInput("Yeni şifrelerin uyumu"))
				} else if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if control == "NotFound" {
					writeResponse(w, notFindRecordError())
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

// LostPasswordChange ss
func LostPasswordChange(newPass string, newPassAgain string, email string) (bool, string) {
	person := &Person{}
	err := connection.Collection("users").FindOne(bson.M{"user_infos.user_mail": email}, person)
	if err != nil {
		return false, "NotFound"
	}
	if person.UserInfos.UserPassword == newPass {
		return false, "OldPassAgain"
	}
	if newPass != newPassAgain {
		return false, "PassMatch"
	}
	person.UserInfos.UserPassword = newPass
	errors := connection.Collection("users").Save(person)
	if errors != nil {
		return false, "Save"
	}
	return true, ""

}
