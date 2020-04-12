package main

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
)

func updateProfilePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		var update, control = controlInfo(r.FormValue("name"), r.FormValue("surname"), r.FormValue("email"), r.FormValue("phone"), r.FormValue("password"), r.FormValue("id"))
		if r.FormValue("name") == "" {
			writeResponse(w, requiredInputError("İsim"))
		} else if r.FormValue("surname") == "" {
			writeResponse(w, requiredInputError("Soyisim"))
		} else if r.FormValue("email") == "" {
			writeResponse(w, requiredInputError("E-posta"))
		} else if r.FormValue("phone") == "" {
			writeResponse(w, requiredInputError("Telefon"))
		} else if r.FormValue("password") == "" {
			writeResponse(w, requiredInputError("Şifre"))
		} else if r.FormValue("id") == "" {
			writeResponse(w, requiredInputError("Kullanıcı Numarası"))
		} else {
			if update == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "Save" {
					writeResponse(w, dataBaseSaveError())

				} else if control == "Nil" {
					writeResponse(w, failedRecordError())

				} else if control == "MailOrPhone" {
					writeResponse(w, alreadyDefinedError("Mail veya Telefon numarası"))

				} else if control == "ID" {
					writeResponse(w, objectIDError())

				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func controlInfo(name string, surname string, mail string, phone string, password string, id string) (bool, string) {
	conroltID, errID := checkObjID(id)
	person := &Person{}
	var checkmail = true
	var checkphone = true
	if errID == true {
		err := connection.Collection("users").FindById(bson.ObjectIdHex(conroltID), person)
		if err != nil {
			return false, "Nil"
		}
		getMail := person.UserInfos.UserMail
		getPhone := person.Contacts.UserPhone
		if getMail != mail {
			checkmail = checkMail(mail)
		}
		if getPhone != phone {
			checkphone = checkPhone(phone)
		}
		if checkmail == true && checkphone == true {
			err := connection.Collection("users").FindById(bson.ObjectIdHex(conroltID), person)
			if err != nil {
			}
			person.Contacts.UserRealName = name
			person.Contacts.UserSurname = surname
			person.UserInfos.UserMail = mail
			person.Contacts.UserPhone = phone
			person.UserInfos.UserPassword = password
			arr := connection.Collection("users").Save(person)
			if arr != nil {
				return false, "Save"
			}
			return true, ""
		}
		return false, "MailOrPhone"
	}
	return false, "ID"
}
