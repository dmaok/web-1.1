package apiserver

import (
	"database/sql"
	"github.com/dmaok/web-1.1/internal/app/store/sqlstore"
	"log"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)

	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	store := sqlstore.New(db)
	server := newServer(store)

	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
