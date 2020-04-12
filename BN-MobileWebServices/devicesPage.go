package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func devicesPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.FormValue("userId") == "" {
			writeResponse(w, requiredInputError("Kullanıcı numarası"))
		} else {
			var devices, control = getMyDevices(r.FormValue("userId"))
			if devices == nil {
				if control == "NotFound" {
					writeResponse(w, notFindRecordError())

				} else if control == "ID" {
					writeResponse(w, objectIDError())
				} else {
					writeResponse(w, someThingWentWrong())
				}
			} else {
				writeResponse(w, string(devices))
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func getMyDevices(getID string) ([]byte, string) {
	var data []byte
	var l []*MyDevices
	var user *MyDevices
	id, errID := checkObjID(getID)
	if errID == true {
		beacon := &Beacon{}
		beacons := connection.Collection("beacons").Find(bson.M{"user.user_id": bson.ObjectIdHex(id)})
		for beacons.Next(beacon) {
			beaconTypeConverter := checkBeaconType(beacon.Information.BeaconType)
			user = &MyDevices{beacon.Id, beacon.Information.BeaconName, beaconTypeConverter}
			l = append(l, user)
		}
		data, _ = json.Marshal(l)
		if l == nil {
			return nil, "NotFound"
		}
		if l != nil {
			response := &PersonDevices{l}
			data, _ = json.Marshal(response)
			return addError(data), ""
		}
	}
	return data, "ID"
}
