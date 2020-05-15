package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
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
	var l []*OrdersInWeb
	var orders *OrdersInWeb
	order := &Orders{}
	id, errID := checkObjID(userID)
	if errID == true {
		Order := connection.Collection("orders").Find(bson.M{"contact_info.user_id": bson.ObjectIdHex(id)})
		for Order.Next(order) {
			orders = &OrdersInWeb{order.Created, order.OrderStatus, order.InOrder, order.PaymentType, order.TotalPrice, order.ContactInfo.UserSurname, order.ContactInfo.UserRealName, order.ContactInfo.UserAddress, order.ContactInfo.UserPhone, order.ContactInfo.UserMail}
			l = append(l, orders)
		}
		data, _ = json.Marshal(l)
		if l == nil {
			return nil, "NotFound"
		}
		if l != nil {
			response := &OrdersInWebArr{l}
			data, _ = json.Marshal(response)
			return addError(data), ""
		}
	}
	return data, "ID"

}
