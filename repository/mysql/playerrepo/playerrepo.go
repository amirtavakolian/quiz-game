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
