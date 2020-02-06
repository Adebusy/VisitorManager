package AppCode

import (
	"../messageentities"
	"github.com/jinzhu/gorm"
)

//CheckOfficeAlreadyExist to check of exist
func CheckOfficeAlreadyExist(officeRequest messageentities.CreateOffice) bool {

	var resp bool = false
	return resp
}

//Allmethods interface method
type Allmethods interface {
	CreateOffice(officeRequest messageentities.CreateOffice)
	var period = flag
	var uio = io.Writer

	

}

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
