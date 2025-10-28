package gormprofilerepo

import (
	"github.com/amirtavakolian/quiz-game/param/profileparams"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Profile struct {
	connection *gorm.DB
}

func NewProfileRepo(connection *gorm.DB) Profile {
	return Profile{connection: connection}
}

func (p Profile) Update(profileRepo profileparams.UpdateProfile) error {
	result := p.connection.
		Table("profiles").
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "player_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"fullname", "bio"}),
		}).Create(&profileRepo)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
