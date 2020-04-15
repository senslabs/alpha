package datastore

import (
	"container/ring"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
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
	logger.Error("No db connections available")
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
		return "5432"
	}
	return port
}

func GetConnectionObsolete() *sqlx.DB {
	once.Do(func() {
		r = ring.New(10)
		n := r.Len()

		pgurl := fmt.Sprintf("postgresql://postgres@Sens1234%s:%s/postgres?ssl=false&sslmode=disable", GetCockroachHost(), GetCockroachPort())
		for i := 0; i < n; i++ {
			db, err := sqlx.Connect("postgres", pgurl)
			r.Value = Connection{i, db, err}
			r = r.Next()
		}
	})

	return getNextConnection()
}

var db *sql.DB = nil

func init() {
	initdb()
}

func initdb() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(err)
		}
	}()
	pgurl := fmt.Sprintf("postgresql://postgres:Sens1234@%s:%s/postgres?sslmode=disable", GetCockroachHost(), GetCockroachPort())
	db, err = sql.Open("postgres", pgurl)
	errors.Pie(err)
}

func GetConnection() *sql.DB {
	for {
		if err := db.Ping(); err != nil {
			logger.Error(err)
			logger.Error("DB connection failure... Waiting for sometime before retrying")
			time.Sleep(5 * time.Second)
			initdb()
		} else {
			break
		}
	}
	return db
}

//This is used while writing into DB
func ConvertFieldValue(column string, v interface{}, typeMap map[string]string) interface{} {
	t := typeMap[column]
	switch t {
	case "*datastore.RawMessage":
		return types.Marshal(v)
	}
	return v
}

func getValue(column string, v interface{}, typeMap map[string]string) interface{} {
	v = *(v.(*interface{}))
	t := typeMap[column]

	switch t {
	case "*string":
	case "*int64":
	case "*bool":
		return v
	case "*uuid.UUID":
		return fmt.Sprintf("%s", v)
	case "*datastore.RawMessage":
		if v != nil {
			return types.UnmarshalMap(v.([]byte))
		}
	}
	return v
}

func RowsToMap(r *sql.Rows, reverseFieldMap map[string]string, typeMap map[string]string) []map[string]interface{} {
	columns, err := r.Columns()
	errors.Pie(err)

	var result []map[string]interface{}
	for r.Next() {
		temp := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for i := range temp {
			values[i] = &temp[i]
		}

		err = r.Scan(values...)
		errors.Pie(err)

		m := make(map[string]interface{})
		for i, c := range columns {
			if field, ok := reverseFieldMap[c]; ok {
				m[field] = getValue(field, values[i], typeMap)
			}
		}
		result = append(result, m)
	}
	return result
}

func RowsToMapReflect(rows *sql.Rows) []map[string]interface{} {
	columns, err := rows.Columns()
	errors.Pie(err)

	var result []map[string]interface{}
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		m := map[string]interface{}{}
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			m[col] = fmt.Sprintf("%s", v)
		}
		result = append(result, m)
	}
	return result
}

func TRACE(seq int, msg string) {
	_, f, _, _ := runtime.Caller(1)
	logger.Debugf("%d: [%s] :: %s in (%s)", seq, time.Now(), msg, filepath.Base(f))
}
