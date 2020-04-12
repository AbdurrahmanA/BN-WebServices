package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func lostDevicesPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.FormValue("userId") == "" {
			writeResponse(w, requiredInputError("Kullanıcı numarası"))
		} else {
			var devices, control = getLostDeviceList(r.FormValue("userId"))
			if devices == nil {
				if control == "nil" {
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
func getLostDeviceList(id string) ([]byte, string) {
	var data []byte
	var beaconInfo *LostBeaconInApp
	controlID, errID := checkObjID(id)
	if errID == true {
		var l []*LostBeaconInApp
		beacon := &LostBeacon{}
		lostBeacon := connection.Collection("lost_beacons").Find(bson.M{"user_infos.user_id": bson.ObjectIdHex(controlID)})
		for lostBeacon.Next(beacon) {
			beaconInfo = &LostBeaconInApp{beacon.Created, beacon.LostLat, beacon.LostLong}
			l = append(l, beaconInfo)
		}
		if l == nil {
			return nil, "nil"
		}
		if l != nil {
			response := &Devices{l}
			data, _ = json.Marshal(response)
			return addError(data), ""
		}
	}
	return data, "ID"

}
