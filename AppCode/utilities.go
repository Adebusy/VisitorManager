package AppCode

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	os.Setenv("password", "tylent")
	os.Setenv("dbname", "VisitorDB")
	os.Setenv("server", "10.0.41.101")
	os.Setenv("databaseSchema", "mysql")
	os.Setenv("root", "root")
	os.Setenv(key, "VisitorDB")
	return os.Getenv(key)
}

var db *sql.DB
var err error
var server = "sterlingazuredb.database.windows.net"
var port = 1433
var user = "dbuser"
var password = "Sterling123"
var database = "aanswebdb"

//DatabaseConnectionString connection string to database
var DatabaseConnectionString = "root:Adebusy100@localhost:3306/VisitorDB"

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

//InitDBConnection connection
func InitDBConnection() {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	} else {
		fmt.Println("no eerror")
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("no eerror 2")
	}
}

//Checkerr checking error
func Checkerr(err error) {
	if err != nil {
		panic(err.Error)
	}
}
