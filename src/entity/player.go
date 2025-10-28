package entity

type Player struct {
	ID          int     `json:"id" gorm:"column:id"`
	Name        *string `json:"name" gorm:"column:firstname"`
	Family      *string `json:"family" gorm:"column:lastname"`
	PhoneNumber string  `json:"phone_number" gorm:"column:phone_number"`
}
