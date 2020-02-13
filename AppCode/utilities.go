package AppCode

import (
	"os"
	"regexp"

	"github.com/Adebusy/VisitorsManager/messageentities"
)

//ValidateEmail used to validate email address
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

func getVariables(key string) string {
	os.Setenv(key, "")
	os.Setenv(key, "tylent")
	os.Setenv(key, "VisitorDB")
	os.Setenv(key, "10.0.41.101")

	os.Setenv(key, "mysql")
	os.Setenv(key, "root")
	os.Setenv(key, "password")

	os.Setenv(key, "VisitorDB")
	return os.Getenv(key)
}

//DbDriver Database driver
const DbDriver = "mysql"

//User database user
const User = "root"

//Password password to db
const Password = "password"

//DbName database name
const DbName = "VisitorDB"

//DatabaseConnectionString connection string to database
const DatabaseConnectionString = "root:root@tcp(http://localhost:8443)/VisitorDB?charset=utf8"

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
