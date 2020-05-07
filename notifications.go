package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//NotificationsForAllUsers bütün kullanıcılara bildirim göndermek için gerekli yapı
type NotificationsForAllUsers struct {
	AppID            string   `json:"app_id"`
	IncludedSegments []string `json:"included_segments"`
	Data             struct {
		Foo string `json:"foo"`
	} `json:"data"`
	Contents struct {
		En string `json:"en"`
	} `json:"contents"`
	Heading struct {
		En string `json:"en"`
	} `json:"headings"`
}

//NotificationsForUserID userid ile bildirim göndermek için gerekli yapı
type NotificationsForUserID struct {
	AppID            string   `json:"app_id"`
	IncludePlayerIds []string `json:"include_player_ids"`
	Data             struct {
		Foo string `json:"foo"`
	} `json:"data"`
	Contents struct {
		En string `json:"en"`
	} `json:"contents"`
	Heading struct {
		En string `json:"en"`
	} `json:"headings"`
}

//AppID ss
var AppID = "56914a6b-2697-4b79-a35c-0ecc952047c3"
var client = http.Client{}

func (msg NotificationsForAllUsers) pushNotificationAllUsers(msgForUsers string, title string) bool {
	var s []string
	s = append(s, "All")

	msg.AppID = AppID
	msg.Data.Foo = "bar"
	msg.IncludedSegments = s
	msg.Contents.En = msgForUsers
	msg.Heading.En = title

	out, err := json.Marshal(msg)
	if err != nil {
		return false
	}

	reqBody := strings.NewReader(string(out))
	request, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", reqBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic NDBjM2I0YTMtNDNkNS00NTgwLWE2MWYtOGNkY2MxNzUyYTdk")
	if err != nil {
		return false
	}
	resp, err := client.Do(request)
	if err != nil {
		return false
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
	if resp.StatusCode != 200 {
		return false
	}
	return true

}
func (msg NotificationsForUserID) pushNotificationPlayerID(ID []string, msgForUsers string, title string) bool {

	msg.AppID = AppID
	msg.Data.Foo = "bar"
	msg.IncludePlayerIds = ID
	msg.Contents.En = msgForUsers
	msg.Heading.En = title

	out, err := json.Marshal(msg)
	if err != nil {
		return false
	}
	reqBody := strings.NewReader(string(out))
	request, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", reqBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic NDBjM2I0YTMtNDNkNS00NTgwLWE2MWYtOGNkY2MxNzUyYTdk")
	if err != nil {
		return false
	}
	resp, err := client.Do(request)
	if err != nil {
		return false
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
	if resp.StatusCode != 200 {
		return false
	}
	return true
}
