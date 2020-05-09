package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func notificationIDListPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("token") == "" {
			writeResponse(w, requiredInputError("token"))
		} else {
			var users, str = notificationList(r.FormValue("token"))
			if users == nil {
				if str == "Token" {
					writeResponse(w, invalidPermission())
				} else {
					writeResponse(w, someThingWentWrong())
				}
			} else {
				writeResponse(w, string(users))
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}
func notificationList(token string) ([]byte, string) {
	controlToken := checkPermission(token)
	var data []byte

	if controlToken == false {
		return data, "Token"
	}
	var l []*NotificationsIDList
	var user *NotificationsIDList
	person := &Person{}
	persons := connection.Collection("users").Find(bson.M{"push_infos.push_id": bson.M{"$ne": ""}})
	for persons.Next(person) {
		user = &NotificationsIDList{person.PushInfos.PushID, person.Id, person.Contacts.UserRealName, person.Contacts.UserSurname}
		l = append(l, user)
	}
	data, _ = json.Marshal(l)
	if l == nil {
		return nil, "NotFound"
	}
	if l != nil {
		response := &NotificationsIDListArr{l}
		data, _ = json.Marshal(response)
		return addError(data), ""
	}
	return data, ""
}
