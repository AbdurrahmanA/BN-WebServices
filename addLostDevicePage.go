package main

import (
	"net/http"
	"strconv"
	"strings"

	creditcard "github.com/durango/go-credit-card"
	"github.com/globalsign/mgo/bson"
)

func addLostDevicePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("phone") == "" {
			writeResponse(w, requiredInputError("Telefon numarası"))
		} else if r.FormValue("lostDesc") == "" {
			writeResponse(w, requiredInputError("Açıklama"))
		} else if r.FormValue("email") == "" {
			writeResponse(w, requiredInputError("Mail"))
		} else if r.FormValue("creditCardNo") == "" {
			writeResponse(w, requiredInputError("Kart numarsı"))
		} else if r.FormValue("creditCardFullName") == "" {
			writeResponse(w, requiredInputError("Kart üzerindeki isim"))
		} else if r.FormValue("creditCardExDate") == "" {
			writeResponse(w, requiredInputError("Kart son kullanma tarihi"))
		} else if r.FormValue("cvv") == "" {
			writeResponse(w, requiredInputError("Güvenlik kodu"))
		} else if r.FormValue("lastSeen") == "" {
			writeResponse(w, requiredInputError("Son görülme"))
		} else if r.FormValue("lostLat") == "" {
			writeResponse(w, requiredInputError("lostLat"))
		} else if r.FormValue("lostLong") == "" {
			writeResponse(w, requiredInputError("lostLong"))
		} else if r.FormValue("beaconID") == "" {
			writeResponse(w, requiredInputError("Cihaz numarası"))
		} else {
			var creatLostDevice, control = addLostDeviceControl(r.FormValue("phone"), r.FormValue("email"), r.FormValue("creditCardNo"), r.FormValue("creditCardFullName"), r.FormValue("creditCardExDate"), r.FormValue("cvv"), r.FormValue("lastSeen"), r.FormValue("lostLat"), r.FormValue("lostLong"), r.FormValue("beaconID"), r.FormValue("lostDesc"))
			if creatLostDevice == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "Card" {
					writeResponse(w, creditCardError())
				} else if control == "NotFound" {
					writeResponse(w, notFindRecordError())
				} else if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if control == "Phone" {
					writeResponse(w, incorrectInput("Phone"))
				} else if control == "Mail" {
					writeResponse(w, incorrectInput("Mail"))
				} else if control == "ID" {
					writeResponse(w, objectIDError())
				} else if control == "Already" {
					writeResponse(w, alreadyDefinedError("BeaconID"))
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}
func addLostDeviceControl(phone string, email string, creditCardNo string, creditCardFullName string, creditCardExDate string, cvv string, lastSeen string, lostLat string, lostLong string, beaconID string, lostDesc string) (bool, string) {
	beacon := &Beacon{}
	fullDate := strings.Split(creditCardExDate, "/")
	card := creditcard.Card{Number: creditCardNo, Cvv: cvv, Month: fullDate[0], Year: fullDate[1]}
	err := card.Validate()
	if err != nil {
		return false, "Card"
	}
	beaconID, errID := checkObjID(beaconID)
	if errID == true {
		lostDeviceControl := &LostBeacon{}
		err := connection.Collection("lost_beacons").FindOne(bson.M{"beacon_infos.beacon_id": bson.ObjectIdHex(beaconID), "lost_status": true}, lostDeviceControl)
		if err != nil {
			return false, "NotFound"
		}
		if lostDeviceControl.BeaconInfos.BeaconID == bson.ObjectIdHex(beaconID) {
			return false, "Already"
		}
		errBeacon := connection.Collection("beacons").FindById(bson.ObjectIdHex(beaconID), beacon)
		if errBeacon != nil {
			return false, "NotFound"
		}
		controlPhone := checkPhoneNumber(phone)
		if controlPhone != true {
			return false, "Phone"
		}
		controlEmail := checkEmailValidity(email)
		if controlEmail != true {
			return false, "Mail"
		}

		floatLostLat, _ := strconv.ParseFloat(lostLat, 64)
		floatLostLong, _ := strconv.ParseFloat(lostLong, 64)

		lostDevice := &LostBeacon{
			LostDate:   lastSeen,
			LostStatus: true,
			LostLat:    floatLostLat,
			LostLong:   floatLostLong,
		}
		lostDevice.LostDesc = lostDesc
		lostDevice.UserInfos.UserMail = email
		lostDevice.UserInfos.UserPhone = phone
		lostDevice.BeaconInfos.BeaconID = beacon.Id
		lostDevice.UserInfos.UserID = beacon.UserInfos.UserID
		errs := connection.Collection("lost_beacons").Save(lostDevice)
		if errs != nil {
			return false, "Save"
		}
		return true, ""
	}
	return false, "ID"
}
