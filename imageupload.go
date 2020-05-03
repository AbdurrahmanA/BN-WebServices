package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func uploadImage(bs64 string, getID string, imgDesc string, imgType int) (bool, string) {
	image, controlImgBase64 := base64ToImage(bs64)
	if controlImgBase64 == false {
		return false, "Base64"
	}
	saveControl, pathImg := saveImage(image, getID, imgDesc, imgType)
	if saveControl == false {
		if pathImg == "File" {
			return false, "File"
		} else if pathImg == "Remove" {
			return false, "Remove"
		} else if pathImg == "Create" {
			return false, "Create"
		} else if pathImg == "Write" {
			return false, "Write"
		} else if pathImg == "Sync" {
			return false, "Sync"
		}
	}
	return true, pathImg
}

func base64ToImage(bs64 string) ([]byte, bool) {
	dec, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		return nil, false
	}
	return dec, true
}

func saveImage(img []byte, ID string, imgDesc string, imgType int) (bool, string) {

	pathImg := ""
	if imgType == 0 {
		pathImg = "users-images/" + ID + "." + imgDesc
	}
	if imgType == 1 {
		pathImg = "beacons-images/" + ID + "." + imgDesc
	}
	controlFile, err := fileExists(pathImg)
	if err != nil {
		fmt.Println(err.Error())
		return false, "File"
	}
	if controlFile == true {
		err := os.Remove(pathImg)
		if err != nil {
			return false, "Remove"
		}
	}
	f, err := os.Create(pathImg)
	if err != nil {
		return false, "Create"
	}
	if _, err := f.Write(img); err != nil {
		return false, "Write"
	}
	if err := f.Sync(); err != nil {
		return false, "Sync"
	}
	defer f.Close()
	return true, pathImg
}
