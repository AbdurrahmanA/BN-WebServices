package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/globalsign/mgo/bson"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func checkMail(newValues string) bool {
	person := &Person{}
	connection.Collection("users").FindOne(bson.M{"user_infos.user_mail": newValues}, person)
	if person.UserInfos.UserMail != "" {
		return false
	}
	return true
}

func checkPhone(newValues string) bool {
	person := &Person{}
	connection.Collection("users").FindOne(bson.M{"contact_infos.user_phone": newValues}, person)
	if person.Contacts.UserPhone != "" {
		return false
	}
	return true
}
func checkBeaconType(beaconType int) string {
	if beaconType == 1 {
		return "Tasma"
	}
	if beaconType == 2 {
		return "Bileklik"
	}
	if beaconType == 3 {
		return "Anahtarlık"
	}
	if beaconType == 4 {
		return "Kalemlik"
	}
	return ""
}
func checkObjID(id string) (string, bool) {
	var s = bson.IsObjectIdHex(id)
	if s == true {
		return id, true
	}
	return "", false
}
func writeResponse(w http.ResponseWriter, jsonValue string) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(jsonValue))
}
func addError(byteJSON []byte) []byte {
	var m map[string]interface{}
	json.Unmarshal(byteJSON, &m)
	m["error"] = false
	newData, _ := json.Marshal(m)
	return newData
}
func checkPermission(token string) bool {
	person := &Person{}
	connection.Collection("users").FindOne(bson.M{"user_infos.user_web_token": token}, person)
	if person.UserInfos.RoleLvl == 5 {
		return true
	}
	return false
}

func checkPhoneNumber(number string) bool {
	regex := regexp.MustCompile("^[+]?(?:[0-9]{2})?[0-9]{10}$")
	match := regex.MatchString(number)
	if match == true {
		return true
	}
	return false
}
func checkEmailValidity(email string) bool {
	regex := regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-zA-Z0-9-.]+$")

	match := regex.MatchString(email)
	if match == true {
		return true
	}
	return false
}
func sendRegisterMail(token string, email string) bool {
	url := "http://213.14.182.224:8090/registercontrol?token="

	temp := registerEmailTemplate(url + token)

	fromEmail := "abdurrahman262@hotmail.com"
	from := mail.NewEmail("BenimkiNerede", fromEmail)
	subject := "Email Onay"
	to := mail.NewEmail(email, email)
	plainTextContent := "text/html"
	//htmlContent := "<strong>"  "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, temp)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		fmt.Println(response.StatusCode)
		return false
	}
	if response.StatusCode != 202 {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		return false
	}
	return true
}
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
