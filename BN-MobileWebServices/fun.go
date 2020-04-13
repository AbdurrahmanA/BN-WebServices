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
	connection.Collection("users").FindOne(bson.M{"user_info.user_mail": newValues}, person)
	if person.UserInfos.UserMail != "" {
		return false
	}
	return true
}

func checkPhone(newValues string) bool {
	person := &Person{}
	connection.Collection("users").FindOne(bson.M{"contact_info.user_phone": newValues}, person)
	if person.Contacts.UserPhone != "" {
		return false
	}
	return true
}
func checkBeaconType(beaconType int) string {
	if beaconType == 0 {
		return "Tasma"
	}
	if beaconType == 1 {
		return "Bileklik"
	}
	if beaconType == 5 {
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
	connection.Collection("users").FindOne(bson.M{"user_info.user_token": token}, person)
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
	url := "http://92.44.120.164:8090/registercontrol?token="

	var temps = `
	<!DOCTYPE html>
	<html>
	<head>
	</head>
	<meta charset="utf-8">
	<meta http-equiv="x-ua-compatible" content="ie=edge">
	<title>Email Confirmation</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style type="text/css">

	@media screen {
		@font-face {
		font-family: 'Source Sans Pro';
		font-style: normal;
		font-weight: 400;
		src: local('Source Sans Pro Regular'), local('SourceSansPro-Regular'), url(https://fonts.gstatic.com/s/sourcesanspro/v10/ODelI1aHBYDBqgeIAH2zlBM0YzuT7MdOe03otPbuUS0.woff) format('woff');
		}

		@font-face {
		font-family: 'Source Sans Pro';
		font-style: normal;
		font-weight: 700;
		src: local('Source Sans Pro Bold'), local('SourceSansPro-Bold'), url(https://fonts.gstatic.com/s/sourcesanspro/v10/toadOcfmlt9b38dHJxOBGFkQc6VGVFSmCnC_l7QZG60.woff) format('woff');
		}
	}

	/**
	* Avoid browser level font resizing.
	* 1. Windows Mobile
	* 2. iOS / OSX
	*/
	body,
	table,
	td,
	a {
		-ms-text-size-adjust: 100%; /* 1 */
		-webkit-text-size-adjust: 100%; /* 2 */
	}

	/**
	* Remove extra space added to tables and cells in Outlook.
	*/
	table,
	td {
		mso-table-rspace: 0pt;
		mso-table-lspace: 0pt;
	}

	/**
	* Better fluid images in Internet Explorer.
	*/
	img {
		-ms-interpolation-mode: bicubic;
	}

	/**
	* Remove blue links for iOS devices.
	*/
	a[x-apple-data-detectors] {
		font-family: inherit !important;
		font-size: inherit !important;
		font-weight: inherit !important;
		line-height: inherit !important;
		color: inherit !important;
		text-decoration: none !important;
	}

	/**
	* Fix centering issues in Android 4.4.
	*/
	div[style*="margin: 16px 0;"] {
		margin: 0 !important;
	}

	body {
		width: 100% !important;
		height: 100% !important;
		padding: 0 !important;
		margin: 0 !important;
	}

	/**
	* Collapse table borders to avoid space between cells.
	*/
	table {
		border-collapse: collapse !important;
	}

	a {
		color: #1a82e2;
	}

	img {
		height: auto;
		line-height: 100%;
		text-decoration: none;
		border: 0;
		outline: none;
	}
	</style>
	<body style="background-color: #e9ecef;">
	<table border="0" cellpadding="0" cellspacing="0" width="100%">
		<tr>
		<td align="center" bgcolor="#e9ecef">
			<table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
			<tr>
				<td align="center" valign="top" style="padding: 36px 24px;">
				<a href="https://benimkinerede.com" target="_blank" style="display: inline-block;">
					<img src="https://i.hizliresim.com/4CGgwD.png" alt="Logo" border="0" width="48" style="display: block; width: 48px; max-width: 48px; min-width: 48px;">
				</a>
				</td>
			</tr>
			</table>
		</td>
		</tr>
		<tr>
		<td align="center" bgcolor="#e9ecef">
			<table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
			<tr>
				<td align="left" bgcolor="#ffffff" style="padding: 36px 24px 0; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; border-top: 3px solid #d4dadf;">
				<h1 style="margin: 0; font-size: 32px; font-weight: 700; letter-spacing: -1px; line-height: 48px;">Email Onay</h1>
				</td>
			</tr>
			</table>
		</td>
		</tr>
		<tr>
		<td align="center" bgcolor="#e9ecef">
			<table border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width: 600px;">
			<tr>
				<td align="left" bgcolor="#ffffff" style="padding: 24px; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; font-size: 16px; line-height: 24px;">
				<p style="margin: 0;">E-posta adresinizi onaylamak için aşağıdaki düğmeye basın.  <a href="https://benimkinerede.com">BenimkiNerede</a>, ile bir hesap oluşturmadıysanız, bu e-postayı güvenle silebilirsiniz.</p>
				</td>
			</tr>
			<tr>
				<td align="left" bgcolor="#ffffff">
				<table border="0" cellpadding="0" cellspacing="0" width="100%">
					<tr>
					<td align="center" bgcolor="#ffffff" style="padding: 12px;">
						<table border="0" cellpadding="0" cellspacing="0">
						<tr>
							<td align="center" bgcolor="#1a82e2" style="border-radius: 6px;" >
							<a href="` + url + token + `"target="_blank" style="display: inline-block; padding: 16px 36px; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; font-size: 16px; color: #ffffff; text-decoration: none; border-radius: 6px;">Onayla</a>
							</td>
						</tr>
						</table>
					</td>
					</tr>
				</table>
				</td>
			</tr>
			<tr>
				<td align="left" bgcolor="#ffffff" style="padding: 24px; font-family: 'Source Sans Pro', Helvetica, Arial, sans-serif; font-size: 16px; line-height: 24px;">
				<p style="margin: 0;">Bu işe yaramazsa, aşağıdaki bağlantıyı kopyalayıp tarayıcınıza yapıştırın:</p>
				<p style="margin: 0;"><a  href="https://benimkinerede.com" target="_blank">` + url + token + `</a></p>
				</td>
			</tr>
			</table>
		</td>
		</tr>
	</table>
	</body>
	</html>
		`

	fromEmail := "abdurrahman262@hotmail.com"
	from := mail.NewEmail("BenimkiNerede", fromEmail)
	subject := "Email Onay"
	to := mail.NewEmail(email, email)
	plainTextContent := "text/html"
	//htmlContent := "<strong>"  "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, temps)
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
