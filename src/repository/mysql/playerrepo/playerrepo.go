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

func (p Player) Store(phoneNumber string) (int64, error) {
	result, err := p.connection.Exec("INSERT IGNORE INTO players (phone_number) VALUES (?)", phoneNumber)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}
