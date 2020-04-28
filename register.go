package main

import (
	"fmt"
	"net/http"
)

func registerPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("email") == "" {
			writeResponse(w, requiredInputError("E-posta"))
		} else if r.FormValue("password") == "" {
			writeResponse(w, requiredInputError("Şifre"))
		} else if r.FormValue("password_again") == "" {
			writeResponse(w, requiredInputError("Şifre tekrar"))
		} else if r.FormValue("name") == "" {
			writeResponse(w, requiredInputError("İsim"))
		} else if r.FormValue("surname") == "" {
			writeResponse(w, requiredInputError("Soy İsim"))
		} else if r.FormValue("phone") == "" {
			writeResponse(w, requiredInputError("Telefon"))
		} else {
			var user, control = register(r.FormValue("email"), r.FormValue("password"), r.FormValue("password_again"), r.FormValue("name"), r.FormValue("surname"), r.FormValue("phone"))
			if user == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "Password" {
					writeResponse(w, incorrectInput("Şifreler"))
				} else if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if control == "SendMail" {
					writeResponse(w, sendMailError())
				} else if control == "Mail" {
					writeResponse(w, incorrectInput("Mail"))
				} else if control == "Phone" {
					writeResponse(w, incorrectInput("Phone"))
				} else if control == "MailData" {
					writeResponse(w, alreadyDefinedError("Mail"))
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func register(userMail string, userPassword string, userPasswordAgain string, name string, surname string, phone string) (bool, string) {
	person := &Person{}

	if userPassword != userPasswordAgain {
		return false, "Password"
	}
	checkMailValid := checkEmailValidity(userMail)
	if checkMailValid == false {
		return false, "Mail"
	}
	checkMail := checkMail(userMail)
	if checkMail == false {
		return false, "MailData"
	}
	checkPhone := checkPhoneNumber(phone)
	if checkPhone == false {
		return false, "Phone"
	}
	tokenReg := tokenGenerator()
	control := sendRegisterMail(tokenReg, userMail)
	if control != true {
		return false, "SendMail"
	}
	person.UserInfos.UserMail = userMail
	person.UserInfos.UserPassword = userPassword
	person.UserInfos.UserWebToken = tokenReg
	person.UserInfos.RoleLvl = 0
	person.Contacts.UserPhone = phone
	person.Contacts.UserRealName = name
	person.Contacts.UserSurname = surname
	errs := connection.Collection("users").Save(person)
	if errs != nil {
		fmt.Println(errs.Error())
		return false, "Save"
	}

	return true, ""
}
