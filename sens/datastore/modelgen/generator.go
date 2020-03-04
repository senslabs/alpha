package sens

import (
	"encoding/json"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 26257
	user     = "root"
	password = "nahinhai"
	dbname   = "postgres"
)

type ModelInfo struct {
	Table string
	Model string
}

func GetModels(db *sqlx.DB, schema string) []ModelInfo {
	models := []ModelInfo{}
	query := `SELECT table_name AS table, replace(initcap(replace(table_name, '_', ' ')), ' ', '') As model FROM information_schema.tables where table_schema = 'public';`
	err := db.Select(&models, query)
	if err != nil {
		log.Fatal(err)
	}
	for i, mi := range models {
		models[i].Model = strings.TrimSuffix(mi.Model, "s")
	}
	return models
}

type FieldInfo struct {
	TableField string `db:"table_field"`
	ModelField string `db:"model_field"`
	Type       string `db:"type"`
	IsNullable string `db:"is_nullable"`
}

func GetFieldType(field FieldInfo) string {
	switch field.Type {
	case "TIMESTAMP":
		if field.IsNullable == "YES" {
			return "NullTime"
		}
		return "time.Time"
	default:
		if field.IsNullable == "YES" {
			return "NullString"
		}
		return "string"
	}
}

func GetTableFields(db *sqlx.DB, schema string, mi ModelInfo) []FieldInfo {
	fields := []FieldInfo{}
	query := fmt.Sprintf(`SELECT column_name AS table_field, replace(initcap(replace(column_name, '_', ' ')), ' ', '') AS model_field, crdb_sql_type AS type, is_nullable FROM information_schema.columns WHERE table_schema='%s' AND table_name = '%s'`, schema, mi.Table)
	err := db.Select(&fields, query)
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range fields {
		fields[i].Type = GetFieldType(f)
	}
	return fields
}

func GenerateStruct(db *sqlx.DB, schema string, mi ModelInfo) string {
	fields := GetTableFields(db, schema, mi)
	members := []string{}
	fieldMap := map[string]string{}
	for _, f := range fields {
		member := fmt.Sprintf("%s %s `db:\"%s\"`", f.ModelField, f.Type, f.TableField)
		members = append(members, member)
		fieldMap[f.ModelField] = f.TableField
	}
	st := fmt.Sprintf(`type %s struct {
		%s
	}
	`, mi.Model, strings.Join(members, "\n    "))

	fm, err := json.Marshal(fieldMap)
	if err != nil {
		log.Fatal(err)
	}
	return st + fmt.Sprintf(`func Get%sFieldMap() map[string]string {
		return map[string]string%s
		}
		`, mi.Model, fm)
}

func GenerateModels(db *sqlx.DB, schema string, models []ModelInfo) {
	sts := []string{}
	for _, mi := range models {
		st := GenerateStruct(db, schema, mi)
		sts = append(sts, st)
		GenerateFunctions(schema, mi)
	}
	os.Mkdir("models", 0777)
	content := []byte(`package models
	import (
		"time"
	)

	var t time.Time
	
	` + strings.Join(sts, "\n\n"))

	content, err := format.Source(content)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("models/models.go", []byte(content), 0666)
}

func GenerateFunctions(schema string, mi ModelInfo) {
	t, err := template.ParseFiles("models/fn/fn.go.tpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(fmt.Sprintf("models/fn/%s.go", mi.Table))
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, mi)
	if err != nil {
		log.Fatal(err)
	}
}

func Generate(schema string) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db := sqlx.MustConnect("postgres", connString)
	defer db.Close()

	models := GetModels(db, schema)
	GenerateModels(db, schema, models)
}
