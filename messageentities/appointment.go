package messageentities

//BookAppointment for requesting visit
type BookAppointment struct {
	VisitorEmail     string `json:"VisitorEmail" gorm:"VisitorEmail"`
	ProposedDate     string `json:"ProposedDate"`
	ProposedDuration string `json:"ProposedDuration" gorm:"ProposedDuration"`
	Purpose          string `json:"Purpose" gorm:"Purpose"`
	StaffEmail       string `json:"StaffEmail" gorm:"StaffEmail"`
	ScheduleType     string `json:"ScheduleType" gorm:"ScheduleType"` // one-off, daily
	AppointmentTime  string `json:"AppointmentTime" gorm:"AppointmentTime"`
	Location         string `json:"Location" gorm:"Location"`
	AppointmentType  string `json:"AppointmentType" gorm:"AppointmentType"`
	Floor            string `json:"Floor" gorm:"Floor"`
	MeetingRoom      string `json:"MeetingRoom" gorm:"MeetingRoom"`
}
