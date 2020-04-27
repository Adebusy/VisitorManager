package Controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Adebusy/VisitorsManager/AppCode"
	ap "github.com/Adebusy/VisitorsManager/AppCode"
	"github.com/Adebusy/VisitorsManager/messageentities"
	"github.com/swaggo/swag/example/celler/httputil"
)

// CreateVisitor godoc
// @Summary creates new student
// @Produce json
// @Param visitor body messageentities.VisitorsRequest true "Create new visitor"
// @Success 200 {object} messageentities.ResponseManager
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /guest/CreateVisitor [post]
func CreateVisitor(ctx *gin.Context) {
	var resp messageentities.ResponseManager
	var visitorRequestBody messageentities.VisitorsRequest
	if err := ctx.ShouldBindJSON(&visitorRequestBody); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		//return
	}
	//check if visitor exist with email address
	if AppCode.ValidateEmail(visitorRequestBody.EmailAddress) == true {
		//check visitor exist
		if ap.CheckVisitorExist(visitorRequestBody, db) {
			resp.ResponseDescription = "Visitor alreasdy created"
			resp.ResponseCode = "01"
		} else {
			if ap.CreateNewVisitor(visitorRequestBody, db) {
				resp.ResponseDescription = "Resquest submitted successfully..."
				resp.ResponseCode = "00"
			} else {
				resp.ResponseDescription = "unable to create visitor at the moment. Please try again later!!"
				resp.ResponseCode = "00"
			}
		}
	} else {
		resp.ResponseDescription = "Invalid Email address supplied. Please confirm and try again."
		resp.ResponseCode = "01"
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetVisitorByEmail godoc
// @Summary fetch visitor by email address
// @Produce json
// @Param guest body messageentities.GetVisitorsByEmail true "get guest details by email request"
// @Success 200 {object} messageentities.VisitorsDetails
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/guest/GetVisitorByEmail/{emailaddress} [get]
func GetVisitorByEmail(ctx *gin.Context) {
	var VisitorsEmail messageentities.GetVisitorsByEmail
	if err := ctx.ShouldBindJSON(&VisitorsEmail); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
	}
	resp := ap.GetVisitorDetailsByEmail(VisitorsEmail, db)
	ctx.JSON(http.StatusOK, resp)
}

//GetVisitorByEmaillocal for local call
func GetVisitorByEmaillocal(getVisitorOBJ messageentities.GetVisitorsByEmail, db *sql.DB) messageentities.VisitorsDetails {
	return ap.GetVisitorDetailsByEmail(getVisitorOBJ, db)
}
