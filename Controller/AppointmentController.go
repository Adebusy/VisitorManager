package Controller

import (
	"../AppCode"
	"../messageentities"
)

//BookAppintment call to book appointment
func BookAppintment(request messageentities.BookAppointment) messageentities.ResponseManager {

	var respMsg messageentities.ResponseManager
	if request.WhomToSee == "" {
		respMsg.ResponseDescription = "Staff does not exist"
		respMsg.ResponseCode = "01"
	}

	if !AppCode.ValidateEmail(request.WhomToSee) {
		respMsg.ResponseDescription = "Staff Email does not exist"
		respMsg.ResponseCode = "01"
	}

	respMsgs := GetVisitorByEmail(request.VisitorEmail)

	if respMsgs.FullName == "" {
		respMsg.ResponseDescription = "Visitor does not exist."
		respMsg.ResponseCode = "01"
	}

	respMsg = AppCode.SubmitAppointmentRequest(request)
	//submit request
	return respMsg
}
