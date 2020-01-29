package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/CreateVisitor", CreateVisitor).Methods("POST")
}

// CreateVisitor method to create CreateVisitor
func CreateVisitor(w http.ResponseWriter, req *http.Request) {

}
