package AppCode

import "github.com/Adebusy/VisitorsManager/messageentities"

//"../messageentities"

//SubmitAppointmentRequest saves request
func SubmitAppointmentRequest(resquest messageentities.BookAppointment) messageentities.ResponseManager {
	var resp messageentities.ResponseManager

	resp = CheckIfStaffExist(resquest)

	if resp.ResponseCode != "00" {
		return resp
	}

	doinsert := SaveAppointmentRequest(resquest)
	if doinsert > 0 {
		resp.ResponseDescription = "Request submitted successfully."
		resp.ResponseCode = "00"
	} else {
		resp.ResponseDescription = "Service is unable to book appointment at the momment. Please try again later."
		resp.ResponseCode = "01"
	}
	return resp
}
