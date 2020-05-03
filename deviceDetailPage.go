package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
)

func deviceDetailsPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.FormValue("deviceId") == "" {
			writeResponse(w, requiredInputError("Cihaz numarası"))
		} else {
			var devices, control = getMyDeviceDetails(r.FormValue("deviceId"))
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
func getMyDeviceDetails(getID string) ([]byte, string) {
	beacon := &Beacon{}
	var data []byte
	conroltID, errID := checkObjID(getID)
	if errID == true {
		err := connection.Collection("beacons").FindById(bson.ObjectIdHex(conroltID), beacon)
		if err != nil {
			return data, "NotFound"
		}
		var beaconType = checkBeaconType(beacon.Information.BeaconType)
		beaconImg := ""
		if beacon.Information.Image != "" {
			dt := time.Now()
			beaconImg = "http://213.14.182.224:8090/" + beacon.Information.Image + "?day=" + dt.Format("01-02-2006") + "?hour=" + dt.Format("15:04:05")
		}
		beacons := &MyDevicesDetail{beacon.Information.BeaconName, beacon.Information.UUID, beaconType, beacon.Information.Variance, beaconImg}
		data, _ = json.Marshal(beacons)
		return addError(data), ""
	}
	return data, "ID"
}
