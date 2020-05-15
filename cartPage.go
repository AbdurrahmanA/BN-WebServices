package main

import (
	"net/http"
	"strconv"
	"strings"

	creditcard "github.com/durango/go-credit-card"
	"github.com/globalsign/mgo/bson"
)

func cartPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("orders") == "" {
			writeResponse(w, requiredInputError("orders "))
		} else if r.FormValue("user_real_name") == "" {
			writeResponse(w, requiredInputError("Kullanıcı Adı"))
		} else if r.FormValue("user_id") == "" {
			writeResponse(w, requiredInputError("Kullanıcı ID "))
		} else if r.FormValue("user_surname") == "" {
			writeResponse(w, requiredInputError("Kullanıcı Soyadı "))
		} else if r.FormValue("user_address") == "" {
			writeResponse(w, requiredInputError("Adres "))
		} else if r.FormValue("user_phone") == "" {
			writeResponse(w, requiredInputError("Telefon "))
		} else if r.FormValue("email") == "" {
			writeResponse(w, requiredInputError("Telefon "))
		} else if r.FormValue("creditCardNo") == "" {
			writeResponse(w, requiredInputError("creditCardNo"))
		} else if r.FormValue("creditCardFullName") == "" {
			writeResponse(w, requiredInputError("creditCardFullName"))
		} else if r.FormValue("creditCardExDate") == "" {
			writeResponse(w, requiredInputError("creditCardExDate"))
		} else if r.FormValue("cvv") == "" {
			writeResponse(w, requiredInputError("cvv"))
		} else if r.FormValue("totalprice") == "" {
			writeResponse(w, requiredInputError("totalprice"))
		} else {
			var control, str = cartPageControl(r.FormValue("orders"), r.FormValue("user_real_name"), r.FormValue("user_id"), r.FormValue("user_surname"), r.FormValue("user_address"), r.FormValue("user_phone"), r.FormValue("email"), r.FormValue("creditCardNo"), r.FormValue("creditCardFullName"), r.FormValue("creditCardExDate"), r.FormValue("cvv"), r.FormValue("totalprice"))
			if str == true {
				writeResponse(w, succesfullyRecordedError())
			} else {
				if control == "Card" {
					writeResponse(w, creditCardError())
				} else if control == "SaveBeacon" {
					writeResponse(w, incorrectInput("SaveBeacon"))
				} else if control == "Save" {
					writeResponse(w, dataBaseSaveError())
				} else if control == "Phone" {
					writeResponse(w, incorrectInput("Phone"))
				} else if control == "Mail" {
					writeResponse(w, incorrectInput("Mail"))
				} else if control == "ID" {
					writeResponse(w, objectIDError())
				} else if control == "totalPrice" {
					writeResponse(w, incorrectInput("totalPrice"))
				} else if control == "Append" {
					writeResponse(w, incorrectInput("Append"))
				} else {
					writeResponse(w, someThingWentWrong())
				}
			}
		}
	} else {
		writeResponse(w, notValidRequestError(r.Method))
	}
}
func cartPageControl(order string, userRealName string, userID string, userSurname string, userAddress string, userPhone string, email string, creditCardNo string, creditCardFullName string, creditCardExDate string, cvv string, totalPrice string) (string, bool) {
	creditCardNoLast := strings.Replace(creditCardNo, " ", "", -1)
	creditCardExDateLast := strings.Replace(creditCardExDate, " ", "", -1)

	fullDate := strings.Split(creditCardExDateLast, "/")
	card := creditcard.Card{Number: creditCardNoLast, Cvv: cvv, Month: fullDate[0], Year: fullDate[1]}
	err := card.Validate()
	if err != nil {
		return "Card", false
	}
	controlEmail := checkEmailValidity(email)
	if controlEmail == false {
		return "Email", false
	}
	controlPhone := checkPhoneNumber(userPhone)
	if controlPhone == false {
		return "Phone", false
	}
	conroltID, errID := checkObjID(userID)
	if errID != true {
		return "ID", false
	}
	totalPriceFloat, err := strconv.ParseFloat(totalPrice, 64)
	if err != nil {
		return "totalPrice", false
	}
	userObjID := bson.ObjectIdHex(conroltID)

	orders := ordersArrayConvert(order)
	ord := ordersStructFilling(orders)
	beacon := &Beacon{}

	for i := 0; i < len(ord); i++ {
		for j := 0; j < 4; j++ {
			if ord[i].ProductType == j {
				var objIDArr []bson.ObjectId
				quantity := ord[i].Quantity
				results := connection.Collection("beacons").Find(bson.M{"beacon_infos.type": ord[i].ProductType})
				count := 0
				for results.Next(beacon) {
					objIDArr = append(objIDArr, beacon.Id)
					count++
				}
				if count >= quantity {
					for t := 0; t < count; t++ {
						device := &Beacon{}
						err := connection.Collection("beacons").FindById(objIDArr[t], device)
						if err != nil {
							return "Append", false
						}
						device.UserInfos.UserID = userObjID
						device.UserInfos.UserMail = email
						device.UserInfos.UserPhone = userPhone
						errors := connection.Collection("beacons").Save(device)
						if errors != nil {
							return "SaveBeacon", false
						}
					}
				}

			}
		}

	}
	Order := &Orders{}
	Order.InOrder = ord
	Order.OrderStatus = 0
	Order.PaymentType = "Cart"
	Order.TotalPrice = totalPriceFloat
	Order.ContactInfo.UserPhone = userPhone
	Order.ContactInfo.UserAddress = userAddress
	Order.ContactInfo.UserID = userObjID
	Order.ContactInfo.UserMail = email
	Order.ContactInfo.UserSurname = userSurname
	Order.ContactInfo.UserRealName = userRealName
	errors := connection.Collection("orders").Save(Order)
	if errors != nil {
		return "Save", false
	}
	return "", true

}
func ordersArrayConvert(orders string) []string {
	fileSlice := strings.Split(orders, ",")
	stringFiles := strings.Join(fileSlice, "|")
	fileSlices := strings.Split(stringFiles, "|")
	return fileSlices
}
func ordersStructFilling(orders []string) []OrderArrayInMongo {
	var ordersStruct []OrderArrayInMongo
	var order OrderArrayInMongo
	count := 0
	for i := 0; i < len(orders)/6; i++ {
		id := bson.ObjectIdHex(orders[count])
		price, _ := strconv.ParseFloat(orders[count+1], 64)
		ordType, _ := strconv.Atoi(orders[count+2])
		quantity, _ := strconv.Atoi(orders[count+3])
		order = OrderArrayInMongo{id, ordType, price, quantity, orders[count+4], orders[count+5]}
		count = count + 6
		ordersStruct = append(ordersStruct, order)
	}
	return ordersStruct
}
