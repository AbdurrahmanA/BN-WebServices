package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
)

func addBeaconPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("token") == "" {
			writeResponse(w, requiredInputError("Token "))
		} else if r.FormValue("major") == "" {
			writeResponse(w, requiredInputError("Major"))
		} else if r.FormValue("minor") == "" {
			writeResponse(w, requiredInputError("Minor"))
		} else if r.FormValue("uuid") == "" {
			writeResponse(w, requiredInputError("UUID"))
		} else if r.FormValue("type") == "" {
			writeResponse(w, requiredInputError("Type"))
		} else {
			var devices, control = addBeacon(r.FormValue("token"), r.FormValue("major"), r.FormValue("minor"), r.FormValue("uuid"), r.FormValue("type"))
			if devices == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "Perm" {
					writeResponse(w, invalidPermission())
				} else if control == "Major" {
					writeResponse(w, incorrectInput("Major"))
				} else if control == "Minor" {
					writeResponse(w, incorrectInput("Minor"))
				} else if control == "Type" {
					writeResponse(w, incorrectInput("Type"))
				} else if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else {
					writeResponse(w, someThingWentWrong())
					fmt.Println(control)
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}
func addBeacon(token string, major string, minor string, uuid string, beaconType string) (bool, string) {

	tokenControl := checkPermission(token)
	if tokenControl == true {
		majorInt, err := strconv.Atoi(major)
		if err != nil {
			return false, "Major"
		}
		minorInt, err := strconv.Atoi(minor)
		if err != nil {
			return false, "Minor"
		}
		beaconTypeInt, err := strconv.Atoi(beaconType)
		if err != nil {
			return false, "Type"
		}
		if beaconTypeInt > 4 || beaconTypeInt < 0 {
			return false, "Type"
		}
		beacon := &Beacon{}
		beacon.Information.BeaconType = beaconTypeInt
		beacon.Information.Major = majorInt
		beacon.Information.Minor = minorInt
		beacon.Information.UUID = uuid
		beacon.Information.LostStatus = false
		beacon.UserInfos.UserID = bson.NewObjectId()
		errs := connection.Collection("beacons").Save(beacon)
		if errs != nil {
			fmt.Println(errs.Error())
			return false, "Save"
		}
		return true, ""
	}
	return false, "Perm"

}
