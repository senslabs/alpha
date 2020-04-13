package datastore

import (
	"container/ring"
	"fmt"
	"os"
	"sync"

	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/sqlx"

	_ "github.com/lib/pq"
)

var err error
var once sync.Once
var r *ring.Ring

type Connection struct {
	serial int
	db     *sqlx.DB
	err    error
}

func getNextConnection() *sqlx.DB {
	n := r.Len()
	for i := 0; i < n; i++ {
		r = r.Next()
		conn := r.Next().Value.(Connection)
		if conn.err == nil {
			logger.Debugf("Returning connection: %d", conn.serial)
			err := conn.db.Ping()
			if err != nil {
				logger.Errorf("Connection: %d failed", i)
				continue
			}
			return conn.db
		}
	}
	return nil
}

func GetCockroachHost() string {
	host := os.Getenv("COCKROACH_HOST")
	if host == "" {
		return "localhost"
	}
	return host
}

func GetCockroachPort() string {
	port := os.Getenv("COCKROACH_PORT")
	if port == "" {
		return "26257"
	}
	return port
}

func GetConnection() *sqlx.DB {
	once.Do(func() {
		r = ring.New(10)
		n := r.Len()

		pgurl := fmt.Sprintf("postgresql://root@%s:%s/postgres?ssl=false&sslmode=disable", GetCockroachHost(), GetCockroachPort())
		for i := 0; i < n; i++ {
			db, err := sqlx.Connect("postgres", pgurl)
			r.Value = Connection{i, db, err}
			r = r.Next()
		}
	})

	return getNextConnection()
}

// func GetConnection() *sqlx.DB {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			logger.Error(err)
// 		}
// 	}()
// 	pgurl := fmt.Sprintf("postgresql://root@%s:%s/postgres?ssl=false&sslmode=disable", GetCockroachHost(), GetCockroachPort())
// 	db := sqlx.MustConnect("postgres", pgurl)
// }
