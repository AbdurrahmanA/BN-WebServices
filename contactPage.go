package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func contactPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("email") == "" {
			writeResponse(w, requiredInputError("Mail"))
		} else if r.FormValue("msg") == "" {
			writeResponse(w, requiredInputError("msg"))
		} else if r.FormValue("name") == "" {
			writeResponse(w, requiredInputError("name"))
		} else if r.FormValue("surname") == "" {
			writeResponse(w, requiredInputError("surname"))
		} else if r.FormValue("title") == "" {
			writeResponse(w, requiredInputError("title"))
		} else {
			var control = sendContactMail(r.FormValue("email"), r.FormValue("msg"), r.FormValue("title"), r.FormValue("name"), r.FormValue("surname"))
			if control == true {
				writeResponse(w, succesfullyError())
			} else {
				writeResponse(w, sendMailError())
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}

func sendContactMail(email string, msg string, title string, name string, surname string) bool {
	control := checkEmailValidity(email)
	if control == false {
		return false
	}
	temp := " Kimden:" + name + " " + surname + " \n Email: " + email + " \n Konu: " + title + "\n Mesaj: " + msg

	fromEmail := "benimneredeki@gmail.com"
	from := mail.NewEmail("BenimkiNerede", fromEmail)
	subject := "İletişim"
	to := mail.NewEmail("benimneredeki@outlook.com", "benimneredeki@outlook.com")
	plainTextContent := "text/html"
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
