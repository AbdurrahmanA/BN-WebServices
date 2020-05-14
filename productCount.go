package main

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func productCountPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var control, str = productCount()
		if control != nil {
			writeResponse(w, string(control))
		} else {
			if str == "NotFound" {
				writeResponse(w, notFindRecordError())
			} else if str == "Perm" {
				writeResponse(w, invalidPermission())
			} else {
				writeResponse(w, someThingWentWrong())
			}
		}

	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func productCount() ([]byte, string) {
	var data []byte
	var m = make(map[int]int)
	for i := 0; i < 4; i++ {
		results := connection.Collection("beacons").Find(bson.M{"beacon_infos.type": i})
		beacons := &Beacon{}
		count := 0
		for results.Next(beacons) {
			count++
		}
		m[i] = count
	}
	data, _ = json.Marshal(m)
	return addError(data), ""

}
