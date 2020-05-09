package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/globalsign/mgo/bson"
)

func notificationsPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("msg") == "" {
			writeResponse(w, requiredInputError("msg"))
		} else if r.FormValue("id") == "" {
			writeResponse(w, requiredInputError("id"))
		} else if r.FormValue("title") == "" {
			writeResponse(w, requiredInputError("title"))
		} else if r.FormValue("type") == "" {
			writeResponse(w, requiredInputError("type"))
		} else if r.FormValue("importanceType") == "" {
			writeResponse(w, requiredInputError("importanceType"))
		} else {
			var control, str = notificationsPageControl(r.FormValue("msg"), r.FormValue("id"), r.FormValue("title"), r.FormValue("type"), r.FormValue("importanceType"), r.FormValue("userId"))
			if control == false {
				if str == "Noti" {
					writeResponse(w, sendNotificationError())
				} else if str == "Convert" {
					writeResponse(w, incorrectInput("Convert"))
				} else if str == "For" {
					writeResponse(w, incorrectInput("append"))
				} else if str == "Type" {
					writeResponse(w, incorrectInput("importanceType"))
				} else if str == "beaconTypeInt" {
					writeResponse(w, incorrectInput("beaconTypeInt"))
				} else if str == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if str == "ID" {
					writeResponse(w, objectIDError())
				} else {
					writeResponse(w, someThingWentWrong())
				}
			} else {
				writeResponse(w, succesfullyNotificationError())
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}
func notificationsPageControl(msg string, id string, title string, beaconType string, importanceType string, userID string) (bool, string) {
	notificationForUser := NotificationsForUserID{}
	var control bool
	var errs bool
	controlID := ""
	if userID != "" {
		controlID, errs = checkObjID(userID)
		if errs == false {
			return false, "ID"
		}
	}
	importanceTypeInt, err := strconv.Atoi(importanceType)
	if err != nil {
		return false, "Convert"
	}
	if importanceTypeInt > 3 || importanceTypeInt < 0 {
		return false, "Type"
	}
	beaconTypeInt, err := strconv.Atoi(beaconType)
	if err != nil {
		return false, "Convert"
	}
	if beaconTypeInt != 6 && beaconTypeInt != 0 && beaconTypeInt != 1 && beaconTypeInt != 2 && beaconTypeInt != 3 {
		return false, "beaconTypeInt"
	}
	ids := strings.Split(id, ",")

	if ids[0] == "All" {
		controlAll, strAll := notificationsAll(msg, title, importanceTypeInt)
		return controlAll, strAll
	}
	if beaconType != "null" && beaconTypeInt != 6 {
		controlGroup, strGroup := notificationsGroup(msg, title, beaconTypeInt, importanceTypeInt)
		return controlGroup, strGroup
	}
	if controlID != "" {
		control = notificationForUser.pushNotificationPlayerID(ids, msg, title)
		if control != false {
			controlIDBson := bson.ObjectIdHex(controlID)
			notiForUser := &NotificationForUser{}
			notiForUser.Description = msg
			notiForUser.Title = title
			notiForUser.UserID = controlIDBson
			notiForUser.ImportanceType = importanceTypeInt
			errs := connection.Collection("notifications").Save(notiForUser)
			if errs != nil {
				return false, "Save"
			}
		}
	}
	return control, "Noti"
}

func notificationsAll(msg string, title string, importanceTypeInt int) (bool, string) {
	notificationForAll := NotificationsForAllUsers{}

	control := notificationForAll.pushNotificationAllUsers(msg, title)
	if control != false {
		notiForAll := &NotificationForAll{}
		notiForAll.Description = msg
		notiForAll.Title = title
		notiForAll.ImportanceType = importanceTypeInt
		errs := connection.Collection("notifications").Save(notiForAll)
		if errs != nil {
			return false, "Save"
		}
	}

	return control, "Noti"
}
func notificationsGroup(msg string, title string, beaconTypeInt int, importanceTypeInt int) (bool, string) {
	person := &Person{}
	beacon := &Beacon{}
	notificationForUsers := NotificationsForUserID{}
	var userIds []string
	pushIds := connection.Collection("beacons").Find(bson.M{"beacon_infos.type": beaconTypeInt})
	for pushIds.Next(beacon) {
		s, _ := bson.ObjectId.MarshalJSON(beacon.UserInfos.UserID)
		t := strings.ReplaceAll(string(s), `"`, "")
		userIds = append(userIds, t)
	}
	var pushNotIds []string
	for i := 0; i < len(userIds); i++ {
		err := connection.Collection("users").FindById(bson.ObjectIdHex(userIds[i]), person)
		if err != nil {
			return false, "For"
		}
		pushNotIds = append(pushNotIds, person.PushInfos.PushID)
	}
	control := notificationForUsers.pushNotificationPlayerID(pushNotIds, msg, title)
	if control != false {
		notiForGroup := &NotificationForGroups{}
		notiForGroup.Description = msg
		notiForGroup.Title = title
		notiForGroup.ImportanceType = importanceTypeInt
		notiForGroup.GroupTypes = beaconTypeInt
		errs := connection.Collection("notifications").Save(notiForGroup)
		if errs != nil {
			return false, "Save"
		}
	}
	return control, "Noti"
}
