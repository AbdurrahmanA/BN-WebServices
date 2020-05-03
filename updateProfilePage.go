package main

import (
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func updateProfilePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" || r.Method == "POST" {
		if r.FormValue("name") == "" {
			writeResponse(w, requiredInputError("İsim"))
		} else if r.FormValue("surname") == "" {
			writeResponse(w, requiredInputError("Soyisim"))
		} else if r.FormValue("phone") == "" {
			writeResponse(w, requiredInputError("Telefon"))
		} else if r.FormValue("img") == "" {
			writeResponse(w, requiredInputError("Image"))
		} else if r.FormValue("imgDesc") == "" {
			writeResponse(w, requiredInputError("Image Description"))
		} else if r.FormValue("id") == "" {
			writeResponse(w, requiredInputError("Kullanıcı Numarası"))
		} else {
			var update, control = updateProfile(r.FormValue("name"), r.FormValue("surname"), r.FormValue("phone"), r.FormValue("img"), r.FormValue("imgDesc"), r.FormValue("id"))
			if update == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if control == "MailOrPhone" {
					writeResponse(w, alreadyDefinedError("Telefon numarası"))
				} else if control == "ID" {
					writeResponse(w, objectIDError())
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
				} else if control == "NotFound" {
					writeResponse(w, notFindRecordError())
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func updateProfile(name string, surname string, phone string, img string, imgDesc string, id string) (bool, string) {
	conroltID, errID := checkObjID(id)
	person := &Person{}

	if errID == true {
		err := connection.Collection("users").FindById(bson.ObjectIdHex(conroltID), person)
		if err != nil {
			return false, "NotFound"
		}
		getPhone := person.Contacts.UserPhone
		var checkphone = true

		if getPhone != phone {
			checkphone = checkPhone(phone)
		}
		if checkphone == false {
			return false, "MailOrPhone"
		}
		imgPath := person.UserInfos.Image
		if img != "null" && imgDesc != "null" {
			imgPathControl := false
			imgPathControl, imgPath = uploadImage(img, conroltID, imgDesc, 0)
			if imgPathControl != true {
				return false, imgPath
			}
		}
		person.Contacts.UserRealName = name
		person.Contacts.UserSurname = surname
		person.Contacts.UserPhone = phone
		person.UserInfos.Image = imgPath
		arr := connection.Collection("users").Save(person)
		if arr != nil {
			fmt.Println("1")
			return false, "Save"
		}
		return true, ""
	}
	return false, "ID"
}
