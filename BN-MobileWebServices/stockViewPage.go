package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func stockViewPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("token") == "" {
			writeResponse(w, requiredInputError("Anahtar"))
		} else {
			var view, control = stockView(r.FormValue("token"))
			if view != nil {
				writeResponse(w, string(view))
			} else {
				if control == "Permission" {
					writeResponse(w, invalidPermission())

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

func stockView(token string) ([]byte, string) {
	var data []byte
	var device *StockView
	control := checkPermission(token)
	if control == false {
		return data, "Permission"
	}
	var l []*StockView
	beacon := &Beacon{}
	beacons := connection.Collection("beacons").Find(bson.M{"user_infos.user_mail": ""})
	for beacons.Next(beacon) {
		device = &StockView{beacon.Information.UUID, beacon.Information.Major, beacon.Information.Minor}
		l = append(l, device)
	}
	data, _ = json.Marshal(l)
	if l == nil {
		return nil, "NotFound"
	}

	response := &StockViewArray{l}
	data, _ = json.Marshal(response)
	return addError(data), ""

}
