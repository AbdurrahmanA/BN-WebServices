package main

import (
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
)

func updateDevicePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		var update, control = controlDeviceInfo(r.FormValue("name"), r.FormValue("type"), r.FormValue("variance"), r.FormValue("beaconID"))
		if r.FormValue("name") == "" {
			writeResponse(w, requiredInputError("İsim"))
		} else if r.FormValue("type") == "" {
			writeResponse(w, requiredInputError("Tip"))
		} else if r.FormValue("variance") == "" {
			writeResponse(w, requiredInputError("Güven aralığı"))
		} else if r.FormValue("beaconID") == "" {
			writeResponse(w, requiredInputError("Cihaz numarası"))
		} else {
			if update == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "ID" {
					writeResponse(w, objectIDError())

				} else if control == "Nil" {
					writeResponse(w, notFindRecordError())

				} else if control == "Save" {
					writeResponse(w, dataBaseSaveError())

				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func controlDeviceInfo(name string, getType string, variance string, beaconID string) (bool, string) {
	device := &Beacon{}
	conroltID, errID := checkObjID(beaconID)
	if errID == true {
		err := connection.Collection("beacons").FindById(bson.ObjectIdHex(conroltID), device)
		if err != nil {
			return false, "Nil"
		}
		newVariance, _ := strconv.Atoi(variance)
		newType, _ := strconv.Atoi(getType)

		device.Information.BeaconName = name
		device.Information.BeaconType = newType
		device.Information.Variance = newVariance
		errors := connection.Collection("beacons").Save(device)
		if errors != nil {
			return false, "Save"
		}
		return true, ""
	}
	return false, "ID"
}
