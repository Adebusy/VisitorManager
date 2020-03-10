package messageentities

//BookAppointment for requesting visit
type BookAppointment struct {
	VisitorEmail     string `json:"VisitorEmail,omitempty"`
	DepartmentID     int    `json:"DepartmentID,omitempty"`
	UnitId           int    `json:"UnitId,omitempty"`
	StaffID          int    `json:"StaffID,omitempty"`
	ProposedDate     string `json:"ProposedDate,omitempty"`
	ProposedDuration string `json:"ProposedDuration,omitempty"`
	Purpose          string `json:"Purpose,omitempty"`
}
