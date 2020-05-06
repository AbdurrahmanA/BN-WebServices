package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("email") == "" {
			writeResponse(w, requiredInputError("E-posta"))
		} else if r.FormValue("password") == "" {
			writeResponse(w, requiredInputError("Şifre"))
		} else if r.FormValue("loginType") == "" {
			writeResponse(w, requiredInputError("loginType"))
		} else {
			var user, control = findUser(r.FormValue("email"), r.FormValue("password"), r.FormValue("loginType"), r.FormValue("pushId"))
			if control == "Login" {
				writeResponse(w, string(user))
			} else if control == "Notfound" {
				writeResponse(w, notFindRecordError())

			} else if control == "PushID" {
				writeResponse(w, requiredInputError("PushID"))

			} else if control == "Lvl" {
				writeResponse(w, invalidLoginRequest())

			} else if control == "Parse" {
				writeResponse(w, incorrectInput("Parse"))

			} else if control == "Mail" {
				writeResponse(w, incorrectInput("Mail"))
			} else {
				writeResponse(w, someThingWentWrong())
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func findUser(userMail string, userPassword string, loginType string, pushID string) ([]byte, string) {
	person := &Person{}
	var data []byte
	controlMail := checkEmailValidity(userMail)
	if controlMail != true {
		return data, "Mail"
	}
	err := connection.Collection("users").FindOne(bson.M{"user_infos.user_mail": userMail, "user_infos.user_password": userPassword}, person)
	if err != nil {
		return data, "Notfound"
	}
	lvl := person.UserInfos.RoleLvl
	if lvl == 0 {
		return data, "Lvl"
	}
	if loginType == "web" {

		person.UserInfos.UserWebToken = tokenGenerator()
		connection.Collection("users").Save(person)
		user := &Userjon{person.UserInfos.UserWebToken}
		data, err = json.Marshal(user)
		if err != nil {
			return data, "Parse"
		}
		return addError(data), "Login"
	} else if loginType == "mobil" {
		if pushID == "" {
			return data, "PushID"
		}
		person.UserInfos.UserMobilToken = tokenGenerator()
		person.PushInfos.PushID = pushID
		connection.Collection("users").Save(person)
		user := &Userjon{person.UserInfos.UserMobilToken}
		data, err = json.Marshal(user)
		if err != nil {
			return data, "Parse"
		}
		return addError(data), "Login"
	} else {
		return data, "loginType"
	}

}
