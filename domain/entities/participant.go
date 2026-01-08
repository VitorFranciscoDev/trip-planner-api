package entities

type Participant struct {
	ID     int    `json:"id"`
	TripID int    `json:"trip_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}
