package messageentities

//ResponseManager to all requests
type ResponseManager struct {
	ResponseDescription string `json:"ResponseDescription,omitempty"`
	ResponseCode        string `json:"ResponseCode,omitempty"`
}

//VisitorsDetails with email address
type VisitorsDetails struct {
	FullName     string `json:"FullName,omitempty"`
	HomeAddress  string `json:"HomeAddress,omitempty"`
	Company      string `json:"Company,omitempty"`
	PhoneNumber  string `json:"PhoneNumber,omitempty"`
	City         string `json:"City,omitempty"`
	State        string `json:"State,omitempty"`
	EmailAddress string `json:"EmailAddress,omitempty"`
	RegNo        string `json:"RegNo,omitempty"`
	RegDate      string `json:"RegDate,omitempty"`
}
