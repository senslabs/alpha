package models

import (
	"container/ring"
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var err error
var once sync.Once
var r *ring.Ring

type Connection struct {
	db  *sqlx.DB
	err error
}

func getNextConnection() *sqlx.DB {
	n := r.Len()
	for i := 0; i < n; i++ {
		r = r.Next()
		conn := r.Next().Value.(Connection)
		if conn.err == nil {
			return conn.db
		}
	}
	return nil
}

func GetConnection() *sqlx.DB {
	once.Do(func() {
		r = ring.New(10)
		n := r.Len()
		host := os.Getenv("COCKROACH_HOST")
		if host == "" {
			host = "localhost"
		}
		pgurl := fmt.Sprintf("postgresql://root@%s:26257/postgres?ssl=false&sslmode=disable", host)
		for i := 0; i < n; i++ {
			db, err := sqlx.Open("postgres", pgurl)
			r.Value = Connection{db, err}
			r = r.Next()
		}
	})

	return getNextConnection()
}
