package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/globalsign/mgo/bson"
)

func uploadImage(bs64 string, getID string) {
	//conroltID := checkObjID(getID)
	beacon := &Beacon{}

	var image = base64ToImage(bs64)
	saveControl, pathImg := saveImage(image, getID)
	result := connection.Collection("beacons").FindOne(bson.M{"infos.uuid": getID}, beacon)
	if result != nil {
		fmt.Println(result.Error())
	} else {
		if saveControl == true {
			beacon.Information.Image = pathImg
			connection.Collection("beacons").Save(beacon)
		}
	}

}

func base64ToImage(bs64 string) []byte {
	dec, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		panic(err)
	}

	return dec
}

func saveImage(img []byte, ID string) (bool, string) {

	loc, _ := time.LoadLocation("Asia/Istanbul")
	now := time.Now().In(loc).Format("2006_01_02_15_04_05")

	var pathImg = "../beacons-images/" + ID + "_" + now + ".jpg"
	f, err := os.Create(pathImg)
	if err != nil {
		fmt.Println(err.Error())
		return false, "false"
	}
	defer f.Close()

	if _, err := f.Write(img); err != nil {
		fmt.Println(err.Error())
		return false, "false"

	}
	if err := f.Sync(); err != nil {

		fmt.Println(err.Error())
		return false, "false"

	}
	pathImg = "beacons-images/" + ID + "_" + now + ".jpg"
	return true, pathImg
}
