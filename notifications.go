package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func pushNotification() {
	client := http.Client{}
	reqBody := strings.NewReader(`{
		"app_id": "56914a6b-2697-4b79-a35c-0ecc952047c3",
		"included_segments": "All",
		"contents":{"en":"asdasd"}
	}`)

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
