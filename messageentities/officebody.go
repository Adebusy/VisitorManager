package messageentities

import "time"

//CreateOffice office object
type CreateOffice struct {
	OfficeName        string `json:"OfficeName,omitempty"`
	OfficeAddress     string `json:"OfficeAddress,omitempty"`
	OfficeDepartment  string `json:"OfficeDepartment,omitempty"`
	OfficeUnit        string `json:"OfficeUnit,omitempty"`
	OfficePhoneNumber string `json:"OfficePhoneNumber,omitempty"`
	Floor             string `json:"Floor,omitempty"`
}

//OfficeDetails office details
type OfficeDetails struct {
	OfficeName        string    `json:"OfficeName,omitempty"`
	OfficeAddress     string    `json:"OfficeAddress,omitempty"`
	OfficeDepartment  string    `json:"OfficeDepartment,omitempty"`
	OfficeUnit        string    `json:"OfficeUnit,omitempty"`
	OfficePhoneNumber string    `json:"OfficePhoneNumber,omitempty"`
	Floor             string    `json:"Floor,omitempty"`
	OfficeID          int32     `json:"OfficeID,omitempty"`
	DateCreated       time.Time `json:"DateCreated,omitempty"`
}
