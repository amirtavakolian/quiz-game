package profileparams

type UpdateProfile struct {
	Bio         string `json:"bio,omitempty"`
	Fullname    string `json:"fullname,omitempty"`
	PhoneNumber string `json:"phone_number"`
	PlayerID    float64  `json:"player_id"`
}
