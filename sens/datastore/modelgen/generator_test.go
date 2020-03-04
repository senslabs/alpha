package sens

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestGetModels(t *testing.T) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db := sqlx.MustConnect("postgres", connString)
	defer db.Close()

	fmt.Printf("%#v", GetModels(db, "public"))
}

func TestGenerate(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db := sqlx.MustConnect("postgres", connString)
	defer db.Close()
	Generate("public")
}
