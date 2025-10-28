package entity
import "time"

type Profile struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Fullname  string    `gorm:"type:varchar(255);collate:utf8mb4_persian_ci"`
	Avatar    string    `gorm:"type:varchar(255);collate:utf8mb4_persian_ci"`
	Bio       string    `gorm:"type:text;collate:utf8mb4_persian_ci"`
	PlayerID  uint64    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}