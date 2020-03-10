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
	"github.com/joho/godotenv"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {

	godotenv.Load()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		AppCode.GoDotEnvVariable("Server"), AppCode.GoDotEnvVariable("user"), AppCode.GoDotEnvVariable("Password"), AppCode.GoDotEnvVariable("Port"), AppCode.GoDotEnvVariable("Database"))
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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", checkIfAmUp).Methods("GET")
	router.HandleFunc("/CreateVisitor", CreateVisitor).Methods("POST")         //done b=kb 1024
	router.HandleFunc("/GetVisitorByEmail", GetVisitorByEmail).Methods("POST") //done
	router.HandleFunc("/CreateOffice", CreateOffice).Methods("POST")           //done
	router.HandleFunc("/BookAppointment", Appointment).Methods("POST")
	http.ListenAndServe(":8081", router)
}

func checkIfAmUp(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("i am up and running.")
}

// CreateVisitor method to create new Visitor
func CreateVisitor(w http.ResponseWriter, req *http.Request) {
	var visitorRequestBody messageentities.VisitorsRequest
	json.NewDecoder(req.Body).Decode(&visitorRequestBody)
	fmt.Sprintln(visitorRequestBody)
	fmt.Printf("my email address %s ", visitorRequestBody.City)

	fmt.Printf(" my email addressss : %s", visitorRequestBody.Company)
	CreateVisitorsresponse := Controller.CreateVisitor(visitorRequestBody, db)
	_ = CreateVisitorsresponse.ResponseCode
	json.NewEncoder(w).Encode(CreateVisitorsresponse)
}

//GetVisitorByEmail used to get visitors details by email addgo run go ress
func GetVisitorByEmail(w http.ResponseWriter, req *http.Request) {
	var visitorEmail messageentities.GetVisitorsByEmail
	var resp messageentities.ResponseManager
	json.NewDecoder(req.Body).Decode(&visitorEmail)
	if visitorEmail.VisitorEmail == "" {
		resp.ResponseDescription = "Email is required"
		resp.ResponseCode = "01"
		json.NewEncoder(w).Encode(resp)
		return
	}

	fmt.Println(visitorEmail.VisitorEmail)
	if AppCode.ValidateEmail(visitorEmail.VisitorEmail) == false {
		resp.ResponseDescription = "Invalid Email is supplied"
		resp.ResponseCode = "01"
		json.NewEncoder(w).Encode(resp)
		return
	}
	//func GetVisitorByEmail(VisitorsEmail string) messageentities.VisitorsDetails {
	json.NewEncoder(w).Encode(Controller.GetVisitorByEmail(visitorEmail, db))
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
	var officereq messageentities.BookAppointment
	json.NewDecoder(req.Body).Decode(&officereq)

	fmt.Printf("first contact %d", officereq.DepartmentID)
	//validate request body
	responseVal := Controller.BookAppintment(officereq, db)
	json.NewEncoder(w).Encode(responseVal)
}
