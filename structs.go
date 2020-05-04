package main

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

//Person  Kullanıcı verilerine ulaşmamızı sağlayan yapı
type Person struct {
	bongo.DocumentBase `bson:",inline"`
	Contacts           struct {
		UserRealName string `bson:"user_real_name" json:"user_real_name"`
		UserSurname  string `bson:"user_surname" json:"user_surname"`
		UserAddress  string `bson:"user_address" json:"user_address"`
		UserPhone    string `bson:"user_phone" json:"user_phone"`
	} `bson:"contact_infos"  `
	UserInfos struct {
		UserPassword   string `bson:"user_password" json:"user_password"`
		UserMail       string `bson:"user_mail" json:"user_mail"`
		UserWebToken   string `bson:"user_web_token" json:"user_web_token"`
		UserMobilToken string `bson:"user_mobile_token" json:"user_mobile_token"`
		RoleLvl        int    `bson:"role_lvl" json:"role_lvl"`
		Image          string `bson:"img" json:"img"`
	} `bson:"user_infos"  `
}

//Userjon  Giriş işlemi için gerekli dönüşleri oluşturmamızı sağlayan yapı
type Userjon struct {
	UserToken string `json:"user_token" `
}

//LostDevices cihazları dizi döndürmeyi sağlayan yapı
type LostDevices struct {
	Beacons []*LostBeaconInApp `json:"beacons" `
}

//PersonDevices cihazları dizi döndürmeyi sağlayan yapı
type PersonDevices struct {
	Beacons []*MyDevices `json:"beacons" `
}

//Beacon verilerine ulaşmamızı sağlayan yapı
type Beacon struct {
	bongo.DocumentBase `bson:",inline"`
	Information        struct {
		BeaconName string    `bson:"beacon_name" json:"beacon_name"`
		UUID       string    `bson:"uuid" json:"uuid"`
		Major      int       `bson:"major" json:"major"`
		Minor      int       `bson:"minor" json:"minor"`
		Variance   int       `bson:"variance" json:"variance"`
		Image      string    `bson:"image" json:"image"`
		BeaconType int       `bson:"type" json:"type"`
		LastSeen   time.Time `bson:"last_seen" json:"last_seen"`
	} `bson:"beacon_infos"  `
	UserInfos struct {
		UserID    bson.ObjectId `bson:"user_id" json:"user_id" `
		UserMail  string        `bson:"user_mail" json:"user_mail"`
		UserPhone string        `bson:"user_phone" json:"user_phone"`
	} `bson:"user_infos"  `
}

//StockView beacon verileri stock ekranın için hazırlanması
type StockView struct {
	UUID  string ` json:"uuid"`
	Major int    ` json:"major"`
	Minor int    ` json:"minor"`
}

//StockViewArray beacon verileri stock ekranında vermemizi saglar
type StockViewArray struct {
	StockViews []*StockView `json:"stocks" `
}

//Orders Sipariş bilgileri için gerekli veritabanı yapısı
type Orders struct {
	bongo.DocumentBase `bson:",inline"`
	OrderStatus        int     `bson:"order_status" json:"order_status" `
	InOrder            []Order `bson:"orders" json:"orders" `
	PaymentType        string  `bson:"payment_type" json:"payment_type" `
	TotalPrice         int     `bson:"total_price" json:"total_price" `
	ContactInfo        struct {
		UserID      string `bson:"user_id" json:"user_id" `
		UserSurname string `bson:"user_surname" json:"user_surname"`
		UserAddress string `bson:"user_address" json:"user_address"`
		UserPhone   string `bson:"user_phone" json:"user_phone"`
	} `bson:"contact_info" json:"contact_info"`
}

//OrdersInWeb Sipariş bilgileri için gerekli  yapısı
type OrdersInWeb struct {
	OrderStatus int               ` json:"order_status" `
	InOrder     []OrderArrayInWeb ` json:"orders" `
	PaymentType string            ` json:"payment_type" `
	TotalPrice  int               ` json:"total_price" `
	ContactInfo struct {
		UserID      string ` json:"user_id" `
		UserSurname string ` json:"user_surname"`
		UserAddress string ` json:"user_address"`
		UserPhone   string ` json:"user_phone"`
	} ` json:"contact_info"`
}

//OrderArrayInWeb web den gelen ürünlerin düzenlenmesi için gerekli yapı
type OrderArrayInWeb struct {
	ProductID    string  ` json:"product_id" `
	ProductName  string  ` json:"product_name" `
	ProductPrice float32 ` json:"product_price" `
}

//Order Toplam ürünler için gerekli yapı
type Order struct {
	ProductID    bson.ObjectId `bson:"product_id" json:"product_id" `
	ProductName  string        `bson:"product_name" json:"product_name" `
	ProductPrice float32       `bson:"product_price" json:"product_price" `
}

//Log Yapılan işlemlerin takipi için gerekli yapı
type Log struct {
	bongo.DocumentBase `bson:",inline"`
	UserID             bson.ObjectId `bson:"user_id" json:"user_id" `
	ProcessCode        string        `bson:"process_code" json:"process_code" `
	Descripton         string        `bson:"description" json:"description" `
}

//LostBeacon Kayıp beacon verileri için gerekli yapı
type LostBeacon struct {
	bongo.DocumentBase `bson:",inline"`
	LostStatus         bool    `bson:"lost_status" json:"lost_status" `
	LostDate           string  `bson:"lost_date" json:"lost_date" `
	LostLat            float64 `bson:"lost_lat" json:"lost_lat" `
	LostLong           float64 `bson:"lost_long" json:"lost_long" `
	LostDesc           string  `bson:"lost_description" json:"lost_description" `
	UserInfos          struct {
		UserID    bson.ObjectId `bson:"user_id" json:"user_id" `
		UserMail  string        `bson:"user_mail" json:"user_mail"`
		UserPhone string        `bson:"user_phone" json:"user_phone"`
	} `bson:"user_infos"  `
	BeaconInfos struct {
		BeaconID   bson.ObjectId `bson:"beacon_id" json:"beacon_id" `
		BeaconName string        `bson:"beacon_name" json:"beacon_name"`
		UUID       string        `bson:"uuid" json:"uuid"`
		Major      int           `bson:"major" json:"major"`
		Minor      int           `bson:"minor" json:"minor"`
		Variance   int           `bson:"variance" json:"variance"`
		BeaconType int           `bson:"type" json:"type"`
		LastSeen   time.Time     `bson:"last_seen" json:"last_seen"`
	} `bson:"beacon_infos"  `
}

//Products Ürün verileri için gerekli yapı
type Products struct {
	bongo.DocumentBase `bson:",inline"`
	ProductID          bson.ObjectId ` json:"product_id" `
	ProductDescription string        `bson:"product_description" json:"product_description" `
	ProductName        string        `bson:"product_name" json:"product_name" `
	ProductPrice       float32       `bson:"product_price" json:"product_price" `
	ProductType        int           `bson:"type" json:"type" `
}

//ProductsInApp Ürün verilerini gönderme için gerekli yapı
type ProductsInApp struct {
	ProductID          bson.ObjectId ` json:"product_id" `
	ProductDescription string        `json:"product_description" `
	ProductName        string        `json:"product_name" `
	ProductPrice       float32       ` json:"product_price" `
	ProductType        int           ` json:"type" `
}

//ProductsForAdd Ürün eklemek için gerekli yapı
type ProductsForAdd struct {
	bongo.DocumentBase `bson:",inline"`
	ProductDescription string  `bson:"product_description" json:"product_description" `
	ProductName        string  `bson:"product_name" json:"product_name" `
	ProductPrice       float32 `bson:"product_price" json:"product_price" `
	ProductType        int     `bson:"type" json:"type" `
}

//ProductsArray Ürün verilerini  dizi döndürmeyi sağlayan yapı
type ProductsArray struct {
	Products []*ProductsInApp `json:"products" `
}

//UserInfoInApp uygulamaya aktarılan kullanıcı verileri
type UserInfoInApp struct {
	UserID       bson.ObjectId `json:"user_id" `
	UserRealName string        `json:"user_real_name" `
	UserSurname  string        `json:"user_surname" `
	UserPhone    string        `json:"user_phone" `
	UserPassword string        `json:"user_password" `
	UserMail     string        `json:"user_mail"`
	Image        string        `json:"user_img"`
	UserAddress  string        `json:"user_address"`
	RoleLvl      int           `json:"role_lvl"`
}

//MyDevices kullanıcının cihazlarının bilgisi
type MyDevices struct {
	ID         bson.ObjectId ` json:"beacon_id"`
	UUID       string        ` json:"uuid"`
	BeaconName string        ` json:"beacon_name"`
	BeaconType string        ` json:"type"`
	Variance   int           ` json:"variance"`
	Image      string        ` json:"img"`
}

//MyDevicesDetail cihazın gerekli bilgileri
type MyDevicesDetail struct {
	BeaconName string ` json:"beacon_name"`
	UUID       string ` json:"uuid"`
	BeaconType string ` json:"type"`
	Variance   int    ` json:"variance"`
	Image      string ` json:"img"`
}

//MyDevicesDetailAndInfos cihazın ve kişinin gerekli bilgileri
type MyDevicesDetailAndInfos struct {
	BeaconID   bson.ObjectId ` json:"beacon_id"`
	BeaconName string        ` json:"beacon_name"`
	UserMail   string        ` json:"user_mail"`
	UserPhone  string        ` json:"user_phone"`
}

//LostBeaconInApp kayıp cihaz bilgileri
type LostBeaconInApp struct {
	BeaconID   bson.ObjectId `json:"beacon_id" `
	UserPhone  string        `json:"user_phone" `
	UserMail   string        `json:"user_mail"`
	LostDate   time.Time     `json:"lost_date" `
	LostLat    float64       `json:"lost_lat" `
	LostLong   float64       `json:"lost_long" `
	LostDesc   string        `json:"lost_desc" `
	LostStatus bool          `json:"lost_status" `
}

//FindLostBeacon kayıp cihazların sorgulanamsı için gerekili yapı
type FindLostBeacon struct {
	BeaconName string        ` json:"beacon_name"`
	UUID       string        ` json:"uuid"`
	UserID     bson.ObjectId `json:"user_id" `
	UserPhone  string        `json:"user_phone" `
	UserMail   string        `json:"user_mail"`
	LostStatus bool          `json:"lost_status"`
}
