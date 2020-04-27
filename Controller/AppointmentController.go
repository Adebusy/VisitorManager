package Controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"

	"github.com/Adebusy/VisitorsManager/AppCode"
	"github.com/Adebusy/VisitorsManager/messageentities"
)

var db *sql.DB
var err error

// BookAppintment godoc
// @Summary creates new appointment request
// @Produce json
// @Param guest body messageentities.BookAppointment true "Appointment request"
// @Success 200 {object} messageentities.ResponseManager
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/appointment/BookAppintment [post]
func BookAppintment(ctx *gin.Context) {
	var respMsg messageentities.ResponseManager
	var bkAppoint messageentities.BookAppointment
	//func BookAppintment( ctx *gin.Context   appointmentrequest messageentities.BookAppointment, db *sql.DB) messageentities.ResponseManager {
	if err := ctx.ShouldBindJSON(&bkAppoint); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
	}

	fmt.Printf("email address %s", bkAppoint.VisitorEmail)
	if !AppCode.ValidateEmail(bkAppoint.VisitorEmail) {
		respMsg.ResponseDescription = "Visitors Email is not valid"
		respMsg.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, respMsg)
	}
	if !AppCode.ValidateEmail(bkAppoint.StaffEmail) {
		respMsg.ResponseDescription = "Staff Email is not valid"
		respMsg.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, respMsg)
	}
	var getVisitorOBJ messageentities.GetVisitorsByEmail
	getVisitorOBJ.VisitorEmail = bkAppoint.VisitorEmail
	respMsgs := GetVisitorByEmaillocal(getVisitorOBJ, db)
	if respMsgs.FullName == "" {
		respMsg.ResponseDescription = "Visitor does not exist."
		respMsg.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, respMsg)
	}
	doinsert := AppCode.SaveAppointmentRequest(bkAppoint, db)
	if doinsert == true {
		respMsg.ResponseDescription = "Request submitted successfully."
		respMsg.ResponseCode = "00"
	} else {
		respMsg.ResponseDescription = "Service is unable to book appointment at the momment. Please try again later."
		respMsg.ResponseCode = "01"
	}
	ctx.JSON(http.StatusBadRequest, respMsg)
}
