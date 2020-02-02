package Controller

import (
	"../AppCode"
	"../messageentities"
)

//CreateOffice creates new office
func CreateOffice(officeRequest messageentities.CreateOffice) messageentities.ResponseManager {
	return AppCode.CreateNewOffice(officeRequest)
}
