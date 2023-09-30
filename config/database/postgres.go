package database

import (
	"context"
	"os"
	"sync"

	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
)

var Postgres *pg.DB

var lock = &sync.Mutex{}

func InitPostgreOrm() *pg.DB {
	lock.Lock()
	defer lock.Unlock()

	if Postgres != nil {
		log.Info("POSTGRES : using existing instance")
		return Postgres
	}

	log.Info("POSTGRES : create new instance")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	database := os.Getenv("DATABASE_NAME")

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		log.Info("POSTGRES : config from localhost")
		host = DATABASE_HOST
		port = DATABASE_PORT
		user = DATABASE_USER
		password = DATABASE_PASSWORD
		database = DATABASE_NAME
	}

	db := pg.Connect(&pg.Options{
		Addr:     host + ":" + port,
		User:     user,
		Password: password,
		Database: database,
	})
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	Postgres = db
	log.Info("POSTGRES : instance created")
	return Postgres
}
