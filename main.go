package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Adebusy/VisitorsManager/docs"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	cntr "github.com/Adebusy/VisitorsManager/Controller"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1

var db *sql.DB
var err error

func init() {
	//cr.ConnectMe()
}

func main() {
	docs.SwaggerInfo.Title = "Visitors management system"
	docs.SwaggerInfo.Description = "This is a visitor management API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8060"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/CreateVisitor", cntr.CreateVisitor)
	r.GET("/GetVisitorByEmail/:emailaddress", cntr.GetVisitorByEmail)
	r.POST("/CreateOffice", cntr.CreateOffice)
	r.POST("/BookAppointment", cntr.BookAppintment)
	r.Run(":8060")
	// golang new request squash to test
	//router.HandleFunc("/", checkIfAmUp).Methods("GET")
	// router.HandleFunc("/CreateVisitor", CreateVisitor).Methods("POST")         //done b=kb 1024
	// router.HandleFunc("/GetVisitorByEmail", GetVisitorByEmail).Methods("POST") //done
	// router.HandleFunc("/CreateOffice", CreateOffice).Methods("POST")           //done
	// router.HandleFunc("/BookAppointment", Appointment).Methods("POST")         //done
	// http.ListenAndServe(":8081", router)
}

func checkIfAmUp(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("i am up and running.")
	//go to test squash
}
