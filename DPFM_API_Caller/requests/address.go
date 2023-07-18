package requests

type Address struct {
	AddressID			int     `json:"AddressID"`
	ValidityStartDate	string  `json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	PostalCode			string	`json:"PostalCode"`
	LocalSubRegion		*string	`json:"LocalSubRegion"`
	LocalRegion			string	`json:"LocalRegion"`
	Country				string	`json:"Country"`
	GlobalRegion		string	`json:"GlobalRegion"`
	TimeZone			string	`json:"TimeZone"`
	District			*string	`json:"District"`
	StreetName			string	`json:"StreetName"`
	CityName			string	`json:"CityName"`
	Building			*string	`json:"Building"`
	Floor				*int    `json:"Floor"`
	Room				*int    `json:"Room"`
	CreationDate         string  `json:"CreationDate"`
	LastChangeDate       string  `json:"LastChangeDate"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}
