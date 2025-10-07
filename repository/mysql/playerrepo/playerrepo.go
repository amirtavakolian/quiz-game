package playerrepo

import (
	"database/sql"
)

type Player struct {
	connection *sql.DB
}

func NewPlayerRepo(connection *sql.DB) Player {
	return Player{connection: connection}
}

func (p Player) Store(phoneNumber string) error {
	if _, err := p.connection.Exec("INSERT IGNORE INTO players (phone_number) VALUES (?)", phoneNumber); err != nil {
		return err
	}

	return nil
}
