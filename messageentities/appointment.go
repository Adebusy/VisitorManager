package messageentities

//BookAppointment for requesting visit
type BookAppointment struct {
	VisitorEmail string `json:"VisitorEmail,omitempty"`
	Office       string `json:"Office,omitempty"`
	WhomToSee    string `json:"WhomToSee,omitempty"`
	ProposedTime string `json:"ProposedTime,omitempty"`
	ProposedDate string `json:"ProposedDate,omitempty"`
	Purpose      string `json:"Purpose,omitempty"`
}
