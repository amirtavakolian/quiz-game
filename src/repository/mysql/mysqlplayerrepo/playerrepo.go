package mysqlplayerrepo

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
	result, err := p.connection.Exec("INSERT INTO players (phone_number) VALUES (?) ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)", phoneNumber)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}
	return id, nil
}
