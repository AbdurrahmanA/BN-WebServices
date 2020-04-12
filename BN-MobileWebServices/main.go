package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
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

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8090", handler)

}
