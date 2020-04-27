package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
)

func profilePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("token") == "" {
			writeResponse(w, requiredInputError("Anahtar"))
		} else if r.FormValue("tokenType") == "" {
			writeResponse(w, requiredInputError("tokenType"))
		} else {
			var user, control = getProfileInfos(r.FormValue("token"), r.FormValue("tokenType"))
			if user == nil {
				if control == "Lvl" {
					writeResponse(w, invalidLoginRequest())
				} else if control == "NotFound" {
					writeResponse(w, notFindRecordError())
				} else {
					writeResponse(w, someThingWentWrong())
				}
			} else {
				writeResponse(w, string(user))
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func getProfileInfos(token string, tokenType string) ([]byte, string) {
	var data []byte

	person := &Person{}
	if tokenType == "web" {
		err := connection.Collection("users").FindOne(bson.M{"user_infos.user_web_token": token}, person)
		if err != nil {
			return data, "NotFound"
		}
	} else if tokenType == "mobil" {
		err := connection.Collection("users").FindOne(bson.M{"user_infos.user_mobile_token": token}, person)
		if err != nil {
			return data, "NotFound"
		}
	} else {
		return data, "tokenType"
	}
	lvl := person.UserInfos.RoleLvl
	if lvl == 0 {
		return data, "Lvl"
	}
	personImg := ""
	if person.UserInfos.Image != "" {
		dt := time.Now()
		personImg = "http://213.14.182.224:8090/" + person.UserInfos.Image + "?day=" + dt.Format("01-02-2006") + "?hour=" + dt.Format("15:04:05")
	}
	user := &UserInfoInApp{
		person.Id,
		person.Contacts.UserRealName,
		person.Contacts.UserSurname,
		person.Contacts.UserPhone,
		person.UserInfos.UserPassword,
		person.UserInfos.UserMail,
		personImg,
		person.Contacts.UserAddress,
		person.UserInfos.RoleLvl,
	}
	data, _ = json.Marshal(user)
	return addError(data), ""

}
