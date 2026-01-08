package entities

type LocationSearch struct {
	DisplayName string  `json:"display_name"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
}
