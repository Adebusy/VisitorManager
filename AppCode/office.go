package AppCode

import (
	"strconv"

	"github.com/Adebusy/VisitorsManager/messageentities"
)

//CreateNewOffice creates new office
func CreateNewOffice(officeRequest messageentities.CreateOffice) messageentities.ResponseManager {
	var officeResp messageentities.ResponseManager
	if CheckOfficeAlreadyExist(officeRequest) {
		officeResp.ResponseCode = "01"
		officeResp.ResponseDescription = "Office name already exist"
	}

	//create office
	responseoffice := []string{"officeID", "officeName", "officeCode"}
	responseoffice[0] = "2"
	responseoffice[1] = officeRequest.OfficeName
	responseoffice[2] = officeRequest.OfficeName

	resp, _ := strconv.Atoi(responseoffice[0])

	if resp <= 0 {
		officeResp.ResponseCode = "01"
		officeResp.ResponseDescription = "Unable to create office at the moment. Please try again later."
	} else {
		officeResp.ResponseCode = "00"
		officeResp.ResponseDescription = "officeName:" + responseoffice[1] + "|officeCode:" + responseoffice[2] + "|officeID:" + responseoffice[0]
	}
	return officeResp
}
