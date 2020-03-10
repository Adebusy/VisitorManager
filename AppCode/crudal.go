package AppCode

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Adebusy/VisitorsManager/messageentities"
)

//SaveAppointmentRequest call to create office
// func SaveAppointmentRequest(appointment messageentities.BookAppointment) int32 {
// 	var resp int32 = 0
// 	appnt := tbl_BookAppointments{department_id: appointment.DepartmentID, proposed_date: appointment.ProposedDate, proposed_duration: appointment.ProposedDuration, purpose: appointment.Purpose, staff_id: appointment.StaffID, unit_id: appointment.UnitId, visitor_email: appointment.VisitorEmail}
// 	var connectionstring = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", GoDotEnvVariable("Server"), GoDotEnvVariable("user"), GoDotEnvVariable("Password"), GoDotEnvVariable("Port"), GoDotEnvVariable("Database"))
// 	fmt.Printf("got sss here 1 %s and %d, and %d", appointment.VisitorEmail, appnt.department_id, appnt.department_id)
// 	fmt.Println(connectionstring)
// 	fmt.Printf("got here new ")
// 	db, err := gorm.Open(GoDotEnvVariable("DatabaseServer"), connectionstring)
// 	defer db.Close()
// 	if err == nil {
// 		resp = 1
// 	}
// 	//listofuser := make(map[string]int)

// 	db.NewRecord(appnt)
// 	fmt.Println(appnt)
// 	db.Table("tbl_BookAppointments").Create(&appnt)

// 	db.NewRecord(appnt)
// 	return resp
// }

//SaveAppointmentRequest save using ado.net
func SaveAppointmentRequest(appointment messageentities.BookAppointment, db *sql.DB) bool {
	var resp bool = false
	queryt := fmt.Sprintf("insert into tbl_BookAppointments([VisitorEmail],[ProposedDate],[ProposedDuration],[Purpose],[StaffEmail],[ScheduleType],[AppointmentTime],[Location],[AppointmentType],[Floor],[MeetingRoom]) values('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", appointment.VisitorEmail, appointment.ProposedDate, appointment.ProposedDuration, appointment.Purpose, appointment.StaffEmail, appointment.ScheduleType, appointment.AppointmentTime, appointment.Location, appointment.AppointmentType, appointment.Floor, appointment.MeetingRoom)
	fmt.Println(queryt)
	docprepare, err := db.Prepare(queryt)
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
