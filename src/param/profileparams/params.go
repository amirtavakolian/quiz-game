package profileparams

type UpdateProfile struct {
	Bio         string  `json:"bio,omitempty" gorm:"column:bio"`
	Fullname    string  `json:"fullname,omitempty" gorm:"column:fullname"`
	PhoneNumber string  `json:"phone_number" gorm:"-"`
	PlayerID    float64 `json:"player_id" gorm:"column:player_id"`
}