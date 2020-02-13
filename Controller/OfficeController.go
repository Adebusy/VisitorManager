package Controller

import (
	"github.com/Adebusy/VisitorsManager/AppCode"
	"github.com/Adebusy/VisitorsManager/messageentities"
)

//CreateOffice creates new office
func CreateOffice(officeRequest messageentities.CreateOffice) messageentities.ResponseManager {
	return AppCode.CreateNewOffice(officeRequest)
}
