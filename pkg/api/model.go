package api

type Trip struct {
	Id 				int32		`json:"id,omitempty"`
	Origin 			*City 		`json:"origin,omitempty"`
	Destination 	*City 		`json:"destination,omitempty"`
	Dates 			string 		`json:"dates,omitempty"`
	Price 			float64 	`json:"price,omitempty"`
}

type City struct {
	Id 			int32		`json:"id,omitempty"`
	Name 		string		`json:"name,omitempty"`
}

type TripList struct {
	TripList	[]Trip 		`json:"trips,omitempty"`
	Total		int32		`json:"total,omitempty"`
	Limit		int32		`json:"limit,omitempty"`
}