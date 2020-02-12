package AppCode

import (
	"os"

	"../messageentities"
	"github.com/jinzhu/gorm"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
)

//CheckOfficeAlreadyExist to check of exist
func CheckOfficeAlreadyExist(officeRequest messageentities.CreateOffice) bool {

	var resp bool = false
	return resp
}

//EnvVariable for db
func EnvVariable(key string) string {
	os.Setenv("databaseusername", "VisitorDB")
	return os.Getenv(key)
}

//Allmethods interface method
type Allmethods interface {
	CreateOffice(officeRequest messageentities.CreateOffice)
	//var period = flag
	//var uio = io.Writer

}

//ConnectToDB to mysql database
// func ConnectToDB() {
// 	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8443)/VisitorDB?charset=utf8&parseTime=True")
// 	defer db.Close()
// 	if err == nil {
// 		log.Println("i was able to connect !!!")
// 	}
// }

//CreateOffice call to create office
func CreateOffice(officeRequest messageentities.CreateOffice) int32 {
	var resp int32 = 0
	db, err := gorm.Open("mysql", "root@/VisitorDB?charset=utf8&parseTime=True&loc=Local") //root:""
	defer db.Close()
	if err == nil {
		resp = 1
	}
	newinse := db.Debug().Save(officeRequest)

	if newinse.Error == nil {
		resp = 1
	}

	doinsert := db.Create(officeRequest)

	if doinsert.Error == nil {
		resp = 1
	}
	return resp
}

//SaveAppointmentRequest call to create office
func SaveAppointmentRequest(appointment messageentities.BookAppointment) int32 {
	var resp int32 = 0
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err == nil {
		resp = 1
	}
	newinse := db.Debug().Save(appointment)

	if newinse.Error == nil {
		resp = 1
	}

	doinsert := db.Create(appointment)

	if doinsert.Error == nil {
		resp = 1
	}
	return resp
}
