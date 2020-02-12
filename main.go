package main

import (
	"encoding/json"
	"net/http"

	"./AppCode"
	"./Controller"
	"./messageentities"
	"github.com/gorilla/mux"
)

func main() {

	//connect first
	//AppCode.ConnectToDB()

	router := mux.NewRouter()
	//router.HandleFunc("/", dbconnector).Methods("GET")
	router.HandleFunc("/CreateVisitor", CreateVisitor).Methods("POST")
	router.HandleFunc("/GetVisitorByEmail", GetVisitorByEmail).Methods("GET")
	router.HandleFunc("/CreateOffice", CreateOffice).Methods("POST")
	router.HandleFunc("/BookAppointment", Appointment).Methods("POST")

	http.ListenAndServe(":9061", router)
}

// func dbconnector(w http.ResponseWriter, req *http.Request) {
// 	AppCode.ConnectToDB()
// }

// CreateVisitor method to create new Visitor
func CreateVisitor(w http.ResponseWriter, req *http.Request) {
	var visitorRequestBody messageentities.VisitorsRequest
	json.NewDecoder(req.Body).Decode(&visitorRequestBody)

	CreateVisitorsresponse := Controller.CreateVisitor(visitorRequestBody)
	_ = CreateVisitorsresponse.ResponseCode
	json.NewEncoder(w).Encode(CreateVisitorsresponse)
}

//GetVisitorByEmail used to get visitors details by email address
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
	//validate request body
	responseVal := Controller.CreateOffice(officerequest)
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
