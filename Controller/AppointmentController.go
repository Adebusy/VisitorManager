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
	if appointmentrequest.StaffID == 0 {
		respMsg.ResponseDescription = "Staff does not exist"
		respMsg.ResponseCode = "01"
	}
	fmt.Printf("email address %s", appointmentrequest.VisitorEmail)
	// if !AppCode.ValidateEmail(strconv.Itoa(request.StaffID)) {
	// 	respMsg.ResponseDescription = "Staff Email does not exist"
	// 	respMsg.ResponseCode = "01"
	// }
	// var getVisitorOBJ messageentities.GetVisitorsByEmail
	// getVisitorOBJ.VisitorEmail = request.VisitorEmail
	// respMsgs := GetVisitorByEmail(getVisitorOBJ, db)

	// if respMsgs.FullName == "" {
	// 	respMsg.ResponseDescription = "Visitor does not exist."
	// 	respMsg.ResponseCode = "01"
	// }

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
