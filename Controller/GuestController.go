package Controller

import (
	"github.com/Adebusy/VisitorsManager/messageentities"
)

//CreateVisitor call to create a visitor
func CreateVisitor(messageentities.VisitorsRequest) messageentities.ResponseManager {
	var resp messageentities.ResponseManager
	resp.ResponseDescription = "Resquest submitted successfully..."
	resp.ResponseCode = "00"
	return resp
}

//GetVisitorByEmail fetch visitor email address
func GetVisitorByEmail(VisitorsEmail string) messageentities.VisitorsDetails {
	var visitorResponse messageentities.VisitorsDetails
	_ = VisitorsEmail
	return visitorResponse
}

//VisitRequest for visit
func VisitRequest(visitReq messageentities.VisitorsRequest) messageentities.VisitorsDetails {
	var visitorResponse messageentities.VisitorsDetails
	_ = visitReq
	return visitorResponse

}
