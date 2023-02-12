package postgres

import (
	"database/sql"
	"errors"
	stores2 "ganhaum.henrybarreto.dev/internal/stores"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

var NewErrConnectDatabase = func(err error) error {
	return errors.Join(stores2.ErrConnectDatabase, err)
}

var NewErrPingDatabase = func(err error) error {
	return errors.Join(stores2.ErrPingDatabase, err)
}

func NewStore(username string, password string, database string) (stores2.Stores, error) {
	db, err := sql.Open("postgres", "postgres://"+username+":"+password+"@postgres/"+database+"?sslmode=disable")
	if err != nil {
		return nil, NewErrConnectDatabase(err)
	}

	if err := db.Ping(); err != nil {
		return nil, NewErrPingDatabase(err)
	}

	return &Store{db: db}, nil
}
