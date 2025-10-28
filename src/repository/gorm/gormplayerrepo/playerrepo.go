package gormplayerrepo

import (
	"github.com/amirtavakolian/quiz-game/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Player struct {
	connection *gorm.DB
}

func NewPlayerRepo(connection *gorm.DB) Player {
	return Player{connection: connection}
}

func (p Player) Store(phoneNumber string) (int64, error) {
	if err := p.connection.
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&entity.Player{PhoneNumber: phoneNumber}).Error; err != nil {
		return 0, err
	}

	var player entity.Player
	if err := p.connection.
		Select("id").
		Where("phone_number = ?", phoneNumber).
		First(&player).Error; err != nil {
		return 0, err
	}

	return int64(player.ID), nil
}
