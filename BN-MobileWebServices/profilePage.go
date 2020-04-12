package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func profilePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("token") == "" {
			writeResponse(w, requiredInputError("Anahtar"))
		} else {
			var user = getProfileInfos(r.FormValue("token"))
			if user == nil {
				writeResponse(w, notFindRecordError())
			} else {
				writeResponse(w, string(user))
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func getProfileInfos(token string) []byte {
	var data []byte

	person := &Person{}
	err := connection.Collection("users").FindOne(bson.M{"user_info.user_token": token}, person)
	if err != nil {
		return data
	}
	user := &UserInfoInApp{person.Id, person.Contacts.UserRealName, person.Contacts.UserSurname, person.Contacts.UserPhone, person.UserInfos.UserPassword, person.UserInfos.UserMail}
	data, _ = json.Marshal(user)
	return addError(data)

}
