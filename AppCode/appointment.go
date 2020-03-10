package AppCode

import (
	"database/sql"

	"github.com/Adebusy/VisitorsManager/messageentities"
)

//"../messageentities"

//SubmitAppointmentRequest saves request
func SubmitAppointmentRequest(resquest messageentities.BookAppointment, db *sql.DB) messageentities.ResponseManager {
	var resp messageentities.ResponseManager

	// resp = CheckIfStaffExist(resquest)

	// if resp.ResponseCode != "00" {
	// 	return resp
	// }

	doinsert := SaveAppointmentRequest(resquest, db)
	if doinsert == true {
		resp.ResponseDescription = "Request submitted successfully."
		resp.ResponseCode = "00"
	} else {
		resp.ResponseDescription = "Service is unable to book appointment at the momment. Please try again later."
		resp.ResponseCode = "01"
	}
	return resp
}
