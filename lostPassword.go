package main

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func lostPasswordPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("email") == "" {
			writeResponse(w, requiredInputError("email"))
		} else {
			var control, str = lostPasswordControl(r.FormValue("email"))
			if control == false {
				if str == "Mail" {
					writeResponse(w, incorrectInput("Mail"))
				} else if str == "Notfound" {
					writeResponse(w, notFindRecordError())
				} else if str == "SendMail" {
					writeResponse(w, sendMailError())
				} else if str == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else {
					writeResponse(w, someThingWentWrong())
				}
			} else {
				writeResponse(w, succesfullyError())
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}
func lostPasswordControl(email string) (bool, string) {
	person := &Person{}

	checkMailValid := checkEmailValidity(email)
	if checkMailValid == false {
		return false, "Mail"
	}
	err := connection.Collection("users").FindOne(bson.M{"user_infos.user_mail": email}, person)
	if err != nil {
		return false, "Notfound"
	}
	tokenReg1 := tokenGenerator()
	tokenReg2 := tokenGenerator()

	control := sendLostPasswordMail(tokenReg1+tokenReg2, email)
	if control != true {
		return false, "SendMail"
	}
	person.UserInfos.UserWebToken = tokenReg1 + tokenReg2
	errs := connection.Collection("users").Save(person)
	if errs != nil {
		return false, "Save"
	}
	return true, ""
}
