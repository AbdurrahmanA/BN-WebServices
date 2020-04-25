package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func checkLostDevicePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("uuid") == "" {
			writeResponse(w, requiredInputError("UUID"))
		} else {
			var control, controlStr, devices = checkLostDevice(r.FormValue("uuid"))
			if control == false {
				if controlStr == "Parse" {
					writeResponse(w, incorrectInput("Json Parse"))
				} else if controlStr == "false" {
					writeResponse(w, notFindRecordError())
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

func checkLostDevice(uuid string) (bool, string, []byte) {
	var data []byte
	lostDevice := &LostBeacon{}
	var beaconInfo *FindLostBeacon
	beacons := connection.Collection("lost_beacons").Find(bson.M{"beacon_infos.uuid": uuid})
	for beacons.Next(lostDevice) {
		beaconInfo = &FindLostBeacon{lostDevice.BeaconInfos.BeaconName, lostDevice.BeaconInfos.UUID, lostDevice.UserInfos.UserID, lostDevice.UserInfos.UserPhone, lostDevice.UserInfos.UserMail, lostDevice.LostStatus}
	}
	if beaconInfo.LostStatus == true {
		data, err := json.Marshal(beaconInfo)
		if err != nil {
			return false, "Parse", data
		}
		return true, "", data
	}
	return false, "false", data
}
