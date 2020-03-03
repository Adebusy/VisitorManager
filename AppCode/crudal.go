package AppCode

import (
	"github.com/Adebusy/VisitorsManager/messageentities"
)

//Testselect resting conn
// func Testselect(db *sql.DB) {
// 	fmt.Println("\n==> testselect")
// 	results, errs := db.Query(`select Id, Name from Department`)

// 	if errs != nil {
// 		fmt.Printf("new error message " + errs.Error())
// 	}
// 	Checkerr(errs)
// 	for results.Next() {
// 		var dep messageentities.Department
// 		errs = results.Scan(&dep.Id, &dep.Name)
// 		Checkerr(errs)
// 		fmt.Printf("%d\t%s \n", dep.Id, dep.Name)
// 	}
// }

//EnvVariable for db
// func EnvVariable(key string) string {
// 	os.Setenv("databaseusername", "VisitorDB")
// 	return os.Getenv(key)
// }

//Checkerr for msg

//CreateOffice call to create office
// func CreateOffice(officeRequest messageentities.CreateOffice) int32 {
// 	var resp int32 = 0
// 	db, err := gorm.Open("mysql", "root:Adebusy100@tcp(localhost:3306)/visitorSchema?tls=skip-verify&autocommit=true") //root:""
// 	//db, err := gorm.Open("mysql", "root@/VisitorDB?charset=utf8&parseTime=True&loc=Local") //root:""

// 	defer db.Close()
// 	if err == nil {
// 		resp = 1
// 	}
// 	newinse := db.Debug().Save(officeRequest)

// 	if newinse.Error == nil {
// 		resp = 1
// 	}

// 	doinsert := db.Create(officeRequest)

// 	if doinsert.Error == nil {
// 		resp = 1
// 	}
// 	return resp
// }

//SaveAppointmentRequest call to create office
func SaveAppointmentRequest(appointment messageentities.BookAppointment) int32 {
	var resp int32 = 0
	// // 	//db, err := gorm.Open("mysql", "root:Adebusy100@/visitorSchema?charset=utf8&parseTime=True&loc=Local")
	// // 	//defer db.Close()
	// // 	//if err == nil {
	// // 	//resp = 1
	// // 	//}

	// // 	newinse := db.Debug().Save(appointment)

	// // 	if newinse.Error == nil {
	// // 		resp = 1
	// // 	}

	// // 	doinsert := db.Create(appointment)

	// // 	if doinsert.Error == nil {
	// // 		resp = 1
	// // 	}
	return resp
}
