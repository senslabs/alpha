package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var db *sql.DB
var err error
var once sync.Once

func GetConnection() *sql.DB {
	once.Do(func() {
		pgurl := "postgresql://root@localhost:26257/postgres?ssl=false&sslmode=disable"
		db, err = sql.Open("postgres", pgurl)
		if err != nil {
			log.Fatal(err)
		}
	})
	return db
}
