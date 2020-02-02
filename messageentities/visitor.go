package messageentities

//VisitorsRequest request body
type VisitorsRequest struct {
	FullName     string `json:"FullName,omitempty"`
	HomeAddress  string `json:"HomeAddress,omitempty"`
	Company      string `json:"Company,omitempty"`
	PhoneNumber  string `json:"PhoneNumber,omitempty"`
	City         string `json:"City,omitempty"`
	State        string `json:"State,omitempty"`
	EmailAddress string `json:"EmailAddress,omitempty"`
}
