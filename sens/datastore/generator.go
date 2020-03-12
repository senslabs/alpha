package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/jmoiron/sqlx"
	"github.com/senslabs/alpha/sens/types"

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
	HasId bool
}

func GetModelInfo(db *sqlx.DB, schema string) []*ModelInfo {
	models := []*ModelInfo{}
	query := fmt.Sprintf(`SELECT table_name AS table, replace(initcap(replace(table_name, '_', ' ')), ' ', '') As model FROM information_schema.tables where table_schema = '%s';`, schema)
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
	case "UUID", "STRING":
		if field.IsNullable == "YES" {
			return "NullString"
		}
		return "string"
	case "TIMESTAMP":
		if field.IsNullable == "YES" {
			return "NullTime"
		}
		return "time.Time"
	case "BOOL":
		return "bool"
	default:
		return "RawMessage"
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

//Generate and return struct for one table
func GenerateModel(db *sqlx.DB, schema string, mi *ModelInfo) string {
	fields := GetTableFields(db, schema, *mi)
	members := []string{}
	fieldMap := map[string]string{}
	for _, f := range fields {
		mi.HasId = mi.HasId || f.TableField == "id"
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

//Generate for all tables
func GenerateModels(db *sqlx.DB, schema string, mis []*ModelInfo) {
	ms := []string{}
	for _, mi := range mis {
		m := GenerateModel(db, schema, mi)
		ms = append(ms, m)
	}
	os.Mkdir("generated/models", 0777)
	content := []byte(`package models
	import (
		"time"
	)

	var t time.Time
	
	` + strings.Join(ms, "\n\n"))

	content, err := format.Source(content)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("generated/models/models.go", []byte(content), 0666)
}

//Generate functions for one table
func GenerateFunction(schema string, mi *ModelInfo) {
	t, err := template.ParseFiles("templates/fn.tpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(fmt.Sprintf("generated/models/fn/%s.go", mi.Table))
	if err != nil {
		log.Fatal(err)
	}

	id := ""
	if mi.HasId {
		id = "m.Id = id"
	}
	err = t.Execute(f, types.Map{
		"Table": mi.Table,
		"Model": mi.Model,
		"Id":    id,
	})
	if err != nil {
		log.Fatal(err)
	}
}

//Generate all functions
func GenerateFunctions(db *sqlx.DB, schema string, mis []*ModelInfo) {
	for _, mi := range mis {
		GenerateFunction(schema, mi)
	}
}

func GenerateDb(object string, model string) {
	t, err := template.ParseFiles("templates/db.tpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(fmt.Sprintf("generated/api/db/%s.go", object))
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, types.Map{"Model": model})
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateRest(table string, model string) {
	t, err := template.ParseFiles("templates/rest.tpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(fmt.Sprintf("generated/api/rest/main/%s.go", table))
	if err != nil {
		log.Fatal(err)
	}
	path := strings.ReplaceAll(table, "_", "-")
	err = t.Execute(f, types.Map{
		"Path":  path,
		"Model": model,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateMain(models []string) {
	t, err := template.ParseFiles("templates/main.tpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("generated/api/rest/main/main.go")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, types.Map{
		"Models": models,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Generate(schema string) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db := sqlx.MustConnect("postgres", connString)
	defer db.Close()

	mis := GetModelInfo(db, schema)
	GenerateModels(db, schema, mis)
	GenerateFunctions(db, schema, mis)

	var ms []string
	for _, mi := range mis {
		// GenerateDb(m.Model, m.Model)
		GenerateRest(mi.Table, mi.Model)
		ms = append(ms, mi.Model)
	}
	GenerateMain(ms)
}

func main() {
	schema := os.Args[1]
	Generate(schema)
}
