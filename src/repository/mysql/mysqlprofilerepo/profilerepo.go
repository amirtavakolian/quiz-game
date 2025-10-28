package mysqlprofilerepo

import (
	"database/sql"
	"github.com/amirtavakolian/quiz-game/param/profileparams"
	)

type Profile struct {
	connection *sql.DB
}

func NewProfileRepo(connection *sql.DB) Profile {
	return Profile{connection: connection}
}

func (p Profile) Update(profileRepo profileparams.UpdateProfile) error {
	_, err := p.connection.Exec("INSERT INTO profiles (fullname, bio, player_id) VALUES (?, ?, ?)", profileRepo.Fullname, profileRepo.Bio, profileRepo.PlayerID)

	if err != nil {
		return err
	}
	return nil
}
