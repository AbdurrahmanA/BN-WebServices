package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	connection = conDb()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginPage)
	mux.HandleFunc("/updateprofile", updateProfilePage)
	mux.HandleFunc("/updatedevice", updateDevicePage)
	mux.HandleFunc("/profile", profilePage)
	mux.HandleFunc("/devices", devicesPage)
	mux.HandleFunc("/devicedetail", deviceDetailsPage)
	mux.HandleFunc("/lostdevices", lostDevicesPage)
	mux.HandleFunc("/addlostdevice", addLostDevicePage)
	mux.HandleFunc("/products", productsListPage)
	mux.HandleFunc("/addproduct", addProductPage)
	mux.HandleFunc("/register", registerPage)
	mux.HandleFunc("/registercontrol", validationRegisterPage)
	mux.HandleFunc("/stocks", stockViewPage)
	mux.HandleFunc("/addbeacon", addBeaconPage)
	mux.HandleFunc("/changepassword", passwordChangePage)
	mux.HandleFunc("/checklostdevice", checkLostDevicePage)
	mux.HandleFunc("/cart", cartPage)
	mux.HandleFunc("/notifications", notificationsPage)
	mux.HandleFunc("/mynotifications", myNotificationsListPage)
	mux.HandleFunc("/notificationidlist", notificationIDListPage)
	mux.HandleFunc("/lostpassword", lostPasswordPage)
	mux.HandleFunc("/lostpasswordcontrol", validationLostPasswordPage)
	mux.HandleFunc("/lostpasswordchange", LostPasswordChangePage)
	mux.HandleFunc("/contactmail", contactPage)
	mux.HandleFunc("/deleteitem", deleteItemPage)
	mux.HandleFunc("/updatebeacon", updateBeaconPage)
	mux.HandleFunc("/updateproduct", updateProductPage)
	mux.HandleFunc("/myorders", myOrdersListPage)
	mux.HandleFunc("/productcount", productCountPage)

	mux.Handle("/users-images/", http.StripPrefix("/users-images", http.FileServer(http.Dir("users-images"))))
	mux.Handle("/beacons-images/", http.StripPrefix("/beacons-images", http.FileServer(http.Dir("beacons-images"))))

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8090", handler)

}
