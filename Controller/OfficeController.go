package Controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"

	cr "github.com/Adebusy/VisitorsManager/crudal"
	"github.com/Adebusy/VisitorsManager/messageentities"
)

// CreateOffice godoc
// @Summary create new office
// @Produce json
// @Param office body messageentities.CreateOffice true "office creation request"
// @Success 200 {object} messageentities.ResponseManager
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/office/CreateOffice [post]
func CreateOffice(ctx *gin.Context) {
	var officeResp messageentities.ResponseManager
	var officeRequest messageentities.CreateOffice
	if err := ctx.ShouldBindJSON(&officeRequest); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		//return
	}
	fmt.Printf("email address %s", officeRequest.Floor)
	if cr.CheckIfOfficeAlreadyExist(officeRequest, db) {
		officeResp.ResponseCode = "-1"
		officeResp.ResponseDescription = "Office name already exist"
	} else {
		//create the new office
		if cr.CreateOfficeRequest(officeRequest, db) {
			officeResp.ResponseCode = "00"
			officeResp.ResponseDescription = "officeName:" + officeRequest.OfficeName + "|officeCode:" + officeRequest.OfficePhoneNumber + "|officeID:" + officeRequest.Floor
		} else {
			officeResp.ResponseCode = "01"
			officeResp.ResponseDescription = "Unable to create office at the moment. Please try again later."
		}
	}
	ctx.JSON(http.StatusOK, officeResp)
}
