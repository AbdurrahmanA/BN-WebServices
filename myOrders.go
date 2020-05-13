package main

import (
	"net/http"
)

func myOrdersListPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("userId") == "" {
			writeResponse(w, requiredInputError("userId"))
		} else {
			var control, str = myOrdersList(r.FormValue("userId"))
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

func myOrdersList(userID string) ([]byte, string) {
	var data []byte
	/*var l []*OrdersInWeb
	var notification *OrdersInWeb
	order := &Orders{}
	id, errID := checkObjID(userID)
	if errID == true {
		notifications := connection.Collection("orders").Find(bson.M{"contact_info.user_id": bson.ObjectIdHex(id)})
		for notifications.Next(order) {
			notification = &OrdersInWeb{order.OrderStatus}
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
	}*/
	return data, "ID"

}
