package main

import (
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
)

func updateDevicePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" || r.Method == "POST" {
		if r.FormValue("name") == "" {
			writeResponse(w, requiredInputError("İsim"))
		} else if r.FormValue("variance") == "" {
			writeResponse(w, requiredInputError("Güven aralığı"))
		} else if r.FormValue("img") == "" {
			writeResponse(w, requiredInputError("Image"))
		} else if r.FormValue("imgDesc") == "" {
			writeResponse(w, requiredInputError("Image Description"))
		} else if r.FormValue("beaconID") == "" {
			writeResponse(w, requiredInputError("Cihaz numarası"))
		} else {
			var update, control = controlDeviceInfo(r.FormValue("name"), r.FormValue("variance"), r.FormValue("img"), r.FormValue("imgDesc"), r.FormValue("beaconID"))
			if update == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "ID" {
					writeResponse(w, objectIDError())
				} else if control == "Nil" {
					writeResponse(w, notFindRecordError())
				} else if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if control == "Base64" {
					writeResponse(w, incorrectInput("Base64"))
				} else if control == "File" {
					writeResponse(w, incorrectInput("File"))
				} else if control == "Remove" {
					writeResponse(w, incorrectInput("Remove"))
				} else if control == "Create" {
					writeResponse(w, incorrectInput("Create"))
				} else if control == "Write" {
					writeResponse(w, incorrectInput("Write"))
				} else if control == "Sync" {
					writeResponse(w, incorrectInput("Sync"))
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func controlDeviceInfo(name string, variance string, img string, imgDesc string, beaconID string) (bool, string) {
	device := &Beacon{}
	conroltID, errID := checkObjID(beaconID)

	if errID == true {
		err := connection.Collection("beacons").FindById(bson.ObjectIdHex(conroltID), device)
		if err != nil {
			return false, "Nil"
		}
		imgPath := device.Information.Image
		if img != "null" && imgDesc != "null" {
			imgPathControl := false
			imgPathControl, imgPath = uploadImage(img, conroltID, imgDesc, 1)
			if imgPathControl == false {
				return false, imgPath
			}
		}
		newVariance, _ := strconv.Atoi(variance)
		device.Information.BeaconName = name
		device.Information.Variance = newVariance
		device.Information.Image = imgPath
		errors := connection.Collection("beacons").Save(device)
		if errors != nil {
			return false, "Save"
		}
		return true, ""
	}
	return false, "ID"
}
