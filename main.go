package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Adebusy/VisitorsManager/AppCode"
	"github.com/Adebusy/VisitorsManager/Controller"
	cont "github.com/Adebusy/VisitorsManager/Controller"
	"github.com/Adebusy/VisitorsManager/messageentities"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

var server = "sterlingazuredb.database.windows.net"
var port = 1433
var user = "dbuser"
var password = "Sterling123"
var database = "aanswebdb"

func init() {
	//vairRec := godotenv.Load(`.env`)
	//getusername := goDotEnvVariable("dsad")
	//AppCode.InitDBConnection()
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
	//os.Exit(1)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", checkIfAmUp).Methods("GET")
	router.HandleFunc("/CreateVisitor", CreateVisitor).Methods("POST")
	router.HandleFunc("/GetVisitorByEmail", GetVisitorByEmail).Methods("GET")
	router.HandleFunc("/CreateOffice", CreateOffice).Methods("POST")
	router.HandleFunc("/BookAppointment", Appointment).Methods("POST")

	http.ListenAndServe(":8081", router)
}

func checkIfAmUp(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("i am up and running.")
}

// CreateVisitor method to create new Visitor
func CreateVisitor(w http.ResponseWriter, req *http.Request) {

	//dbConn := AppCode.CheckDatabaseConnection()
	//fmt.Println(dbConn)

	var visitorRequestBody messageentities.VisitorsRequest
	json.NewDecoder(req.Body).Decode(&visitorRequestBody)

	CreateVisitorsresponse := Controller.CreateVisitor(visitorRequestBody)
	_ = CreateVisitorsresponse.ResponseCode
	json.NewEncoder(w).Encode(CreateVisitorsresponse)
}

//GetVisitorByEmail used to get visitors details by email addgo run go ress
func GetVisitorByEmail(w http.ResponseWriter, req *http.Request) {
	var visitorEmail string
	var resp messageentities.ResponseManager
	_ = json.NewDecoder(req.Body).Decode(&visitorEmail)
	if visitorEmail == "" {
		resp.ResponseDescription = "Email is required"
		resp.ResponseCode = "01"
		json.NewEncoder(w).Encode(resp)
	}
	if !AppCode.ValidateEmail(visitorEmail) {
		resp.ResponseDescription = "Invalid Email is supplied"
		resp.ResponseCode = "01"
		json.NewEncoder(w).Encode(resp)
	}
	//func GetVisitorByEmail(VisitorsEmail string) messageentities.VisitorsDetails {
	_ = Controller.GetVisitorByEmail(visitorEmail)

}

//CreateOffice create office for visitor
func CreateOffice(w http.ResponseWriter, req *http.Request) {
	var officerequest messageentities.CreateOffice
	json.NewDecoder(req.Body).Decode(&officerequest)
	responseVal := cont.CreateOffice(officerequest, db)
	json.NewEncoder(w).Encode(responseVal)
}

//Appointment with staff
func Appointment(w http.ResponseWriter, req *http.Request) {
	var officerequest messageentities.BookAppointment
	json.NewDecoder(req.Body).Decode(&officerequest)
	//validate request body
	responseVal := Controller.BookAppintment(officerequest)
	json.NewEncoder(w).Encode(responseVal)
}
