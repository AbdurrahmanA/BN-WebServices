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
		UserRealName string `bson:"user_real_name"`
		UserSurname  string `bson:"user_surname"`
		UserAddress  string `bson:"user_address"`
		UserPhone    string `bson:"user_phone"`
	} `bson:"contact_infos"`
	UserInfos struct {
		UserPassword   string `bson:"user_password"`
		UserMail       string `bson:"user_mail"`
		UserWebToken   string `bson:"user_web_token"`
		UserMobilToken string `bson:"user_mobile_token"`
		RoleLvl        int    `bson:"role_lvl"`
		Image          string `bson:"img" `
	} `bson:"user_infos"`
	PushInfos struct {
		PushID string `bson:"push_id"`
	} `bson:"push_infos"`
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
		LostStatus bool      `bson:"lost_status" json:"lost_status" `
	} `bson:"beacon_infos"  `
	UserInfos struct {
		UserID    bson.ObjectId `bson:"user_id" json:"user_id" `
		UserMail  string        `bson:"user_mail" json:"user_mail"`
		UserPhone string        `bson:"user_phone" json:"user_phone"`
	} `bson:"user_infos"  `
}

//StockView beacon verileri stock ekranın için hazırlanması
type StockView struct {
	UUID  string        ` json:"uuid"`
	Major int           ` json:"major"`
	Minor int           ` json:"minor"`
	Type  string        ` json:"type"`
	ID    bson.ObjectId ` json:"id"`
}

//StockViewArray beacon verileri stock ekranında vermemizi saglar
type StockViewArray struct {
	StockViews []*StockView `json:"stocks" `
}

//Orders Sipariş bilgileri için gerekli veritabanı yapısı
type Orders struct {
	bongo.DocumentBase `bson:",inline"`
	OrderStatus        int                 `bson:"order_status"  `
	InOrder            []OrderArrayInMongo `bson:"orders" `
	PaymentType        string              `bson:"payment_type"  `
	TotalPrice         float64             `bson:"total_price"  `
	ContactInfo        struct {
		UserID       bson.ObjectId `bson:"user_id"  `
		UserSurname  string        `bson:"user_surname" `
		UserRealName string        `bson:"user_real_name" `
		UserAddress  string        `bson:"user_address" `
		UserPhone    string        `bson:"user_phone" `
		UserMail     string        `bson:"user_mail" `
	} `bson:"contact_info" `
}

//OrdersInWeb Sipariş bilgileri için gerekli  yapısı
type OrdersInWeb struct {
	Time         time.Time           ` json:"time" `
	OrderStatus  int                 ` json:"order_status" `
	InOrder      []OrderArrayInMongo ` json:"orders" `
	PaymentType  string              ` json:"payment_type" `
	TotalPrice   float64             ` json:"total_price" `
	UserSurname  string              ` json:"user_surname"`
	UserRealName string              ` json:"user_real_name"`
	UserAddress  string              ` json:"user_address"`
	UserPhone    string              ` json:"user_phone"`
	UserMail     string              ` json:"user_mail"`
}

//OrdersInWebArr Sipariş bilgileri için gerekli  yapısı
type OrdersInWebArr struct {
	Orders []*OrdersInWeb ` json:"orders" `
}

//OrderArrayInMongo Toplam ürünler için gerekli yapı
type OrderArrayInMongo struct {
	ProductID          bson.ObjectId `bson:"product_id" `
	ProductType        int           `bson:"product_type" `
	ProductPrice       float64       `bson:"product_price" `
	Quantity           int           `bson:"quantity" `
	ProductDescription string        `bson:"product_description"`
	ProductName        string        `bson:"product_name"`
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
	ProductDescription string  `bson:"product_description" json:"product_description" `
	ProductName        string  `bson:"product_name" json:"product_name" `
	ProductPrice       float32 `bson:"product_price" json:"product_price" `
	ProductType        int     `bson:"type" json:"type" `
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
	LostStatus bool          ` json:"lost_status"`
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

//NotificationForAll bildirimlerin yapısı
type NotificationForAll struct {
	bongo.DocumentBase `bson:",inline"`
	Title              string `bson:"title"`
	Description        string `bson:"description"`
	ImportanceType     int    `bson:"importance_type"`
}

//NotificationForGroups bildirimlerin yapısı
type NotificationForGroups struct {
	bongo.DocumentBase `bson:",inline"`
	Title              string `bson:"title"`
	GroupTypes         int    `bson:"group_types" `
	Description        string `bson:"description"`
	ImportanceType     int    `bson:"importance_type"`
}

//NotificationForUser bildirimlerin yapısı
type NotificationForUser struct {
	bongo.DocumentBase `bson:",inline"`
	Title              string        `bson:"title"`
	UserID             bson.ObjectId `bson:"user_id"`
	Description        string        `bson:"description"`
	ImportanceType     int           `bson:"importance_type"`
}

//MyNotifications bildirimlerin yapısı
type MyNotifications struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

//MyNotificationsArr bildirimlerin yapısı
type MyNotificationsArr struct {
	Notifications []*MyNotifications `json:"notifications"`
}

//NotificationsIDList bildirimlerin yapısı
type NotificationsIDList struct {
	ID      string        `json:"push_id"`
	ObjID   bson.ObjectId `json:"user_id"`
	Name    string        `json:"name"`
	Surname string        `json:"surname"`
}

//NotificationsIDListArr bildirimlerin yapısı
type NotificationsIDListArr struct {
	Users []*NotificationsIDList `json:"users"`
}
