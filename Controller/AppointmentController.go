package Controller

import (
	"database/sql"
	"fmt"

	"github.com/Adebusy/VisitorsManager/AppCode"
	"github.com/Adebusy/VisitorsManager/messageentities"
)

//BookAppintment call to book appointment
func BookAppintment(appointmentrequest messageentities.BookAppointment, db *sql.DB) messageentities.ResponseManager {
	var respMsg messageentities.ResponseManager

	fmt.Printf("email address %s", appointmentrequest.VisitorEmail)
	if !AppCode.ValidateEmail(appointmentrequest.VisitorEmail) {
		respMsg.ResponseDescription = "Visitors Email is not valid"
		respMsg.ResponseCode = "01"
		return respMsg
	}

	if !AppCode.ValidateEmail(appointmentrequest.StaffEmail) {
		respMsg.ResponseDescription = "Staff Email is not valid"
		respMsg.ResponseCode = "01"
		return respMsg
	}

	var getVisitorOBJ messageentities.GetVisitorsByEmail
	getVisitorOBJ.VisitorEmail = appointmentrequest.VisitorEmail
	respMsgs := GetVisitorByEmail(getVisitorOBJ, db)

	if respMsgs.FullName == "" {
		respMsg.ResponseDescription = "Visitor does not exist."
		respMsg.ResponseCode = "01"
		return respMsg
	}

	doinsert := AppCode.SaveAppointmentRequest(appointmentrequest, db)
	if doinsert == true {
		respMsg.ResponseDescription = "Request submitted successfully."
		respMsg.ResponseCode = "00"
	} else {
		respMsg.ResponseDescription = "Service is unable to book appointment at the momment. Please try again later."
		respMsg.ResponseCode = "01"
	}
	return respMsg
}
