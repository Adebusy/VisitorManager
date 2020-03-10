package AppCode

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Adebusy/VisitorsManager/messageentities"
)

//CheckVisitorExist check if visitor exist
func CheckVisitorExist(visitorRequestBody messageentities.VisitorsRequest, db *sql.DB) bool {
	var vFullName string
	var response bool
	query := fmt.Sprintf(`select FullName from tbl_VisitorsRequest where EmailAddress= '%s'`, visitorRequestBody.EmailAddress)
	resp, err := db.Query(query, visitorRequestBody.EmailAddress)
	if err != nil {
		log.Fatalf("error from sql CheckIfOfficeAlreadyExist " + err.Error())
	} else {
		for resp.Next() {
			err := resp.Scan(&vFullName)
			if err != nil {
				log.Println(err)
			}
			break
		}
		if vFullName != "" {
			response = true
		} else {
			response = false
		}
	}
	return response
}

//CreateNewVisitor to create new visitor
func CreateNewVisitor(RequestBody messageentities.VisitorsRequest, db *sql.DB) bool {
	var resp bool
	dt := time.Now()
	var todaysDate string = strings.Split(dt.String(), " ")[0]
	fmt.Printf("date %s", todaysDate)
	query := fmt.Sprintf(`INSERT INTO tbl_VisitorsRequest  ([FullName],[HomeAddress],[Company],[PhoneNumber],[City],[State],[EmailAddress], [RequestDate]) 
	VALUES ('%s','%s','%s','%s','%s','%s','%s', '%s')`, strings.ToUpper(RequestBody.FullName), strings.ToUpper(RequestBody.HomeAddress), strings.ToUpper(RequestBody.Company), RequestBody.PhoneNumber, strings.ToUpper(RequestBody.City), strings.ToUpper(RequestBody.State), strings.ToUpper(RequestBody.EmailAddress), todaysDate)
	docprepare, err := db.Prepare(query)
	if err != nil {
		log.Panic(err.Error())
	}
	_, errs := docprepare.Exec()
	if errs != nil {
		log.Panic(errs.Error())
		resp = false
	} else {
		resp = true
	}
	return resp
}

//GetVisitorDetailsByEmail get detail
func GetVisitorDetailsByEmail(visitoreEmail messageentities.GetVisitorsByEmail, db *sql.DB) messageentities.VisitorsDetails {
	var resp messageentities.VisitorsDetails
	query := fmt.Sprintf("select FullName,HomeAddress,Company,PhoneNumber,City,State,EmailAddress from tbl_VisitorsRequest where EmailAddress = '%s'", visitoreEmail.VisitorEmail)
	docprepare, err := db.Query(query, visitoreEmail.VisitorEmail)
	if err != nil {
		log.Panic(err.Error())
	}
	for docprepare.Next() {
		respbody := docprepare.Scan(&resp.FullName, &resp.HomeAddress, &resp.Company, &resp.PhoneNumber, &resp.City, &resp.State, &resp.EmailAddress)
		if respbody != nil {
			log.Printf("error sql %s", respbody.Error())
		} else {
			fmt.Println(fmt.Sprintf("GetVisitorDetailsByEmail 3 '%s'", visitoreEmail.VisitorEmail))
		}
		break
	}
	return resp
}
