package entities

type Trip struct {
	ID        int           `json:"id"`
	UserID    int           `json:"user_id"`
	Title     string        `json:"title"`
	Transport string        `json:"transport"`
	Group     []Participant `json:"group"`
	Stops     []Stop        `json:"stops"`
}
