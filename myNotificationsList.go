package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func myNotificationsListPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("userId") == "" {
			writeResponse(w, requiredInputError("userId"))
		} else {
			var control, str = myNotificationsList(r.FormValue("userId"))
			if control != nil {
				writeResponse(w, string(control))
			} else {
				if str == "NotFound" {
					writeResponse(w, notFindRecordError())
				} else if str == "ID" {
					writeResponse(w, incorrectInput("ID"))
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func myNotificationsList(userID string) ([]byte, string) {
	var data []byte
	var l []*MyNotifications
	var notification *MyNotifications
	person := &NotificationForUser{}
	id, errID := checkObjID(userID)
	if errID == true {
		notifications := connection.Collection("notifications").Find(bson.M{"user_id": bson.ObjectIdHex(id)})
		for notifications.Next(person) {
			notification = &MyNotifications{person.Title, person.Description, person.Created}
			l = append(l, notification)
		}
		data, _ = json.Marshal(l)
		if l == nil {
			return nil, "NotFound"
		}
		if l != nil {
			response := &MyNotificationsArr{l}
			data, _ = json.Marshal(response)
			return addError(data), ""
		}
	}
	return data, "ID"
}
