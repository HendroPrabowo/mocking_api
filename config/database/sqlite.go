package database

import (
	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var sqliteDB *gorm.DB

func InitSqlite() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		sqliteDB = db
	})
	return sqliteDB
}
