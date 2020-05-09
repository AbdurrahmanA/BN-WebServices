package main

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func deleteItemPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("id") == "" {
			writeResponse(w, requiredInputError("id"))
		} else if r.FormValue("type") == "" {
			writeResponse(w, requiredInputError("type"))
		} else {
			var control, str = deleteItem(r.FormValue("id"), r.FormValue("type"))
			if control == true {
				writeResponse(w, succesfullyError())
			} else {
				if str == "NotFound" {
					writeResponse(w, notFindRecordError())
				} else if str == "ID" {
					writeResponse(w, objectIDError())
				} else if str == "Type" {
					writeResponse(w, incorrectInput("Type"))
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func deleteItem(id string, ıtemType string) (bool, string) {
	conroltID, errID := checkObjID(id)
	if errID == true {
		if ıtemType == "product" {
			err := connection.Collection("products").DeleteOne(bson.M{"_id": bson.ObjectIdHex(conroltID)})
			if err != nil {
				return false, "NotFound"
			}
			return true, ""
		} else if ıtemType == "beacon" {
			err := connection.Collection("beacons").DeleteOne(bson.M{"_id": bson.ObjectIdHex(conroltID)})
			if err != nil {
				return false, "NotFound"
			}
			return true, ""
		}
		return false, "Type"
	}
	return false, "ID"
}
