package entities

type Stop struct {
	ID          int      `json:"id"`
	TripID      int      `json:"trip_id"`
	Name        string   `json:"name"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Experiences []string `json:"experiences"`
}
