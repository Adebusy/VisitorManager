package Controller

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	crd "github.com/Adebusy/VisitorsManager/crudal"
	"github.com/Adebusy/VisitorsManager/messageentities"
)

//CreateOffice creates new office
func CreateOffice(officeRequest messageentities.CreateOffice, db *sql.DB) messageentities.ResponseManager {
	var officeResp messageentities.ResponseManager
	fmt.Println("request got here 1")
	fmt.Println(officeRequest.OfficeName)
	responseoffice := []string{"officeID", "officeName", "officeCode"}
	if crd.CheckIfOfficeAlreadyExist(officeRequest, db) {
		fmt.Println(`Office name already exist`)
		officeResp.ResponseCode = "01"
		officeResp.ResponseDescription = "Office name already exist"
	} else {
		//create office CreateOfficeRequest
		if crd.CreateOfficeRequest(officeRequest, db) {
			responseoffice[0] = "1"
			responseoffice[1] = officeRequest.OfficeName
			responseoffice[2] = officeRequest.Floor
		} else {
			log.Println(`create  office request returned false for office name `, officeRequest.OfficeName)
			responseoffice[0] = "-1"
			responseoffice[1] = officeRequest.OfficeName
			responseoffice[2] = officeRequest.Floor
		}
	}
	fmt.Printf(`response from create request %s`, responseoffice[0])

	resp, _ := strconv.Atoi(responseoffice[0])
	if resp < 0 {
		officeResp.ResponseCode = "01"
		officeResp.ResponseDescription = "Unable to create office at the moment. Please try again later."
	} else {
		officeResp.ResponseCode = "00"
		officeResp.ResponseDescription = "officeName:" + officeRequest.OfficeName + "|office_Floor:" + officeRequest.Floor
	}
	return officeResp
}
