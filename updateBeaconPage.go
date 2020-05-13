package main

import (
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
)

func updateBeaconPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" || r.Method == "POST" {
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
		} else if r.FormValue("id") == "" {
			writeResponse(w, requiredInputError("id"))
		} else {
			var control, str = updateBeacon(r.FormValue("token"), r.FormValue("major"), r.FormValue("minor"), r.FormValue("uuid"), r.FormValue("type"), r.FormValue("id"))
			if control == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if str == "Perm" {
					writeResponse(w, invalidPermission())
				} else if str == "NotFound" {
					writeResponse(w, notFindRecordError())
				} else if str == "Major" {
					writeResponse(w, incorrectInput("Major"))
				} else if str == "Minor" {
					writeResponse(w, incorrectInput("Minor"))
				} else if str == "Type" {
					writeResponse(w, incorrectInput("Type"))
				} else if str == "Save" {
					writeResponse(w, dataBaseSaveError())
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

func updateBeacon(token string, major string, minor string, uuid string, beaconType string, beaconID string) (bool, string) {
	tokenControl := checkPermission(token)

	if tokenControl == true {
		device := &Beacon{}
		conroltID, errID := checkObjID(beaconID)

		if errID == true {
			err := connection.Collection("beacons").FindById(bson.ObjectIdHex(conroltID), device)
			if err != nil {
				return false, "NotFound"
			}
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
			device.Information.UUID = uuid
			device.Information.Major = majorInt
			device.Information.Minor = minorInt
			device.Information.BeaconType = beaconTypeInt
			errors := connection.Collection("beacons").Save(device)
			if errors != nil {
				return false, "Save"
			}
			return true, ""
		}
		return false, "ID"
	}
	return false, "Perm"
}
