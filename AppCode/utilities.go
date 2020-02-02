package AppCode

import (
	"regexp"

	"../messageentities"
)

//ValidateEmail used to validate email address
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

//ValidateOfficeRequest used to validate office request
func ValidateOfficeRequest(officeRequest messageentities.CreateOffice) messageentities.ResponseManager {
	var resp messageentities.ResponseManager
	if officeRequest.OfficeName == "" {
		resp.ResponseDescription = "OfficeName is required"
		resp.ResponseCode = "01"
	}

	if officeRequest.OfficeAddress == "" {
		resp.ResponseDescription = "OfficeAddress is required"
		resp.ResponseCode = "01"
	}

	if officeRequest.OfficeDepartment == "" {
		resp.ResponseDescription = "OfficeDepartment is required"
		resp.ResponseCode = "01"
	}

	if officeRequest.OfficeUnit == "" {
		resp.ResponseDescription = "OfficeUnit is required"
		resp.ResponseCode = "01"
	}

	if officeRequest.OfficePhoneNumber == "" {
		resp.ResponseDescription = "OfficePhoneNumber is required"
		resp.ResponseCode = "01"
	}

	if officeRequest.Floor == "" {
		resp.ResponseDescription = "Office Floor is required"
		resp.ResponseCode = "01"
	}
	return resp
}
