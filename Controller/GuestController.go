package Controller

import (
	"database/sql"

	"github.com/Adebusy/VisitorsManager/AppCode"
	ap "github.com/Adebusy/VisitorsManager/AppCode"
	"github.com/Adebusy/VisitorsManager/messageentities"
)

//CreateVisitor call to create a visitor
func CreateVisitor(visitorRequestBody messageentities.VisitorsRequest, db *sql.DB) messageentities.ResponseManager {
	var resp messageentities.ResponseManager
	//check if visitor exist with email address
	if AppCode.ValidateEmail(visitorRequestBody.EmailAddress) == true {
		//check visitor exist
		if ap.CheckVisitorExist(visitorRequestBody, db) {
			//create visitor
			resp.ResponseDescription = "Visitor alreasdy created"
			resp.ResponseCode = "01"
		} else {
			if ap.CreateNewVisitor(visitorRequestBody, db) {
				resp.ResponseDescription = "Resquest submitted successfully..."
				resp.ResponseCode = "00"
			} else {
				resp.ResponseDescription = "unable to create visitor at the moment. Please try again later!!"
				resp.ResponseCode = "00"
			}
		}
	} else {
		resp.ResponseDescription = "Invalid Email address supplied. Please confirm and try again."
		resp.ResponseCode = "01"
	}
	return resp
}

//GetVisitorByEmail fetch visitor email address
func GetVisitorByEmail(VisitorsEmail messageentities.GetVisitorsByEmail, db *sql.DB) messageentities.VisitorsDetails {
	return ap.GetVisitorDetailsByEmail(VisitorsEmail, db)
}

//VisitRequest for visit
func VisitRequest(visitReq messageentities.VisitorsRequest) messageentities.VisitorsDetails {
	var visitorResponse messageentities.VisitorsDetails
	_ = visitReq
	return visitorResponse

}
