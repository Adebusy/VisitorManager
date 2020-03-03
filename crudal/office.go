package crudal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Adebusy/VisitorsManager/messageentities"
)

//CheckIfOfficeAlreadyExist confirm office exist
func CheckIfOfficeAlreadyExist(officeRequest messageentities.CreateOffice, db *sql.DB) bool {
	var response bool = false
	var MyOfficeName string = ""
	var oname = strings.ToUpper(officeRequest.OfficeName)
	fmt.Printf(`got here CheckIfOfficeAlreadyExist %s`, oname)
	var myquery = fmt.Sprintf(`select OfficeName from tbl_Office where OfficeName ='%s'`, oname)
	resp, err := db.Query(myquery, oname)
	if err != nil {
		log.Fatalf("error from sql CheckIfOfficeAlreadyExist " + err.Error())
	} else {
		for resp.Next() {
			err := resp.Scan(&MyOfficeName)
			if err != nil {
				log.Println(err)
			}
			break
		}
		if MyOfficeName != "" {
			response = true
		} else {
			response = false
		}
	}
	defer db.Close()
	return response
}

var ctx context.Context
var err error

//CreateOfficeRequest to check
func CreateOfficeRequest(officeRequest messageentities.CreateOffice, db *sql.DB) bool {
	var response bool = false
	query2 := fmt.Sprintf(`insert into tbl_Office(OfficeName,OfficeAddress,OfficeDepartment,
			OfficeUnit,OfficePhoneNumber,OfficeFloor) values ('%s','%s','%s','%s','%s','%s')`, strings.ToUpper(officeRequest.OfficeName), strings.ToUpper(officeRequest.OfficeAddress), strings.ToUpper(officeRequest.OfficeDepartment), strings.ToUpper(officeRequest.OfficeUnit), strings.ToUpper(officeRequest.OfficePhoneNumber), strings.ToUpper(officeRequest.Floor))
	doQuery, errs := db.Prepare(query2)
	if errs != nil {
		panic(errs)
	} else {
		response = true
	}
	_, respError := doQuery.Exec()
	if respError != nil {
		fmt.Printf(respError.Error())
	}
	defer db.Close()
	return response
}
