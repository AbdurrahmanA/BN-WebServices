package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func checkLostDevicePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("uuid") == "" {
			writeResponse(w, requiredInputError("UUID"))
		} else {
			var control, controlStr, devices = checkLostDevice(r.FormValue("uuid"))
			if devices != nil {
				writeResponse(w, string(devices))
			} else {
				if control == false {
					if controlStr == "Parse" {
						writeResponse(w, incorrectInput("Json Parse"))
					} else if controlStr == "NotFound" {
						writeResponse(w, notFindRecordError())
					} else {
						writeResponse(w, someThingWentWrong())
					}
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func checkLostDevice(uuid string) (bool, string, []byte) {
	var data []byte
	lostDevice := &LostBeacon{}
	err := connection.Collection("lost_beacons").FindOne(bson.M{"beacon_infos.uuid": uuid, "lost_status": true}, lostDevice)
	if err != nil {
		fmt.Println(err.Error())
		return false, "NotFound", data
	}
	data, err = json.Marshal(lostDevice)
	if err != nil {
		return false, "Parse", data
	}
	return true, "false", addError(data)
}
