package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//Message ss
type Message struct {
	AppID            string   `json:"app_id"`
	IncludedSegments []string `json:"included_segments"`
	Data             struct {
		Foo string `json:"foo"`
	} `json:"data"`
	Contents struct {
		En string `json:"en"`
	} `json:"contents"`
}

//AppID ss
var AppID = "56914a6b-2697-4b79-a35c-0ecc952047c3"
var s []string

func (msg Message) pushNotificationAllUsers(msgForUsers string) {

	s = append(s, "All")

	client := http.Client{}
	msg.AppID = AppID
	msg.Data.Foo = "bar"
	msg.IncludedSegments = s
	msg.Contents.En = msgForUsers

	out, _ := json.Marshal(msg)

	reqBody := strings.NewReader(string(out))
	request, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", reqBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic NDBjM2I0YTMtNDNkNS00NTgwLWE2MWYtOGNkY2MxNzUyYTdk")
	if err != nil {
		log.Println(err.Error())
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
}
func (msg Message) list() {

}

/*
	reqBody := strings.NewReader(`{
		"app_id": "56914a6b-2697-4b79-a35c-0ecc952047c3",
		"included_segments": ["All",
		"contents":{"en":"asdasd"}
	}`)
*/
