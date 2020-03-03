package AppCode

import (
	"database/sql"

	//"github.com/Adebusy/VisitorsManager/AppCode/crudal"
	"github.com/Adebusy/VisitorsManager/crudal"
	cr "github.com/Adebusy/VisitorsManager/crudal"
	"github.com/Adebusy/VisitorsManager/messageentities"
)

//CreateNewOffice creates new office
func CreateNewOffice(officeRequest messageentities.CreateOffice, db *sql.DB) messageentities.ResponseManager {
	var officeResp messageentities.ResponseManager
	if crudal.CheckIfOfficeAlreadyExist(officeRequest, db) {
		officeResp.ResponseCode = "-1"
		officeResp.ResponseDescription = "Office name already exist"
	} else {
		//create the new office
		if cr.CreateOfficeRequest(officeRequest, db) {
			officeResp.ResponseCode = "00"
			officeResp.ResponseDescription = "officeName:" + officeRequest.OfficeName + "|officeCode:" + officeRequest.OfficePhoneNumber + "|officeID:" + officeRequest.Floor
		} else {
			officeResp.ResponseCode = "01"
			officeResp.ResponseDescription = "Unable to create office at the moment. Please try again later."
		}
	}
	return officeResp
}
