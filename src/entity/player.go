package entity

type Player struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Family      string `json:"family"`
	PhoneNumber string `json:"phone_number"`
}
