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

	"github.com/senslabs/alpha/sens/types"
	"github.com/senslabs/sqlx"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Sens1234"
	dbname   = "postgres"
)

type ModelInfo struct {
	Table string
	Model string
	HasId bool
	Pk    []string
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
	switch strings.ToLower(field.Type) {
	case "string", "text":
		return "*string"
	case "uuid":
		return "*uuid.UUID"
	case "timestamp":
		return "*time.Time"
	case "int8", "integer", "bigint":
		return "*int64"
	case "bool", "boolean":
		return "*bool"
	case "float8", "float", "double precision":
		return "*float64"
	case "json", "jsonb":
		return "*datastore.RawMessage"
	default:
		return "[]byte"
	}
}

func GetTableFields(db *sqlx.DB, schema string, mi ModelInfo) []FieldInfo {
	fields := []FieldInfo{}
	query := fmt.Sprintf(`SELECT column_name AS table_field, replace(initcap(replace(column_name, '_', ' ')), ' ', '') AS model_field, data_type AS type, is_nullable FROM information_schema.columns WHERE table_schema='%s' AND table_name = '%s'`, schema, mi.Table)
	err := db.Select(&fields, query)
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range fields {
		fields[i].Type = GetFieldType(f)
	}
	return fields
}

type FieldConstraint struct {
	TableName      string `db:"table_name"`
	ColumnName     string `db:"column_name"`
	ConstraintName string `db:"constraint_name"`
}

//Update if a model has id or nit
func GetConstraintMap(db *sqlx.DB) map[string][]string {
	var constraints []FieldConstraint
	// query := `SELECT table_name, column_name, constraint_name FROM information_schema.constraint_column_usage WHERE table_catalog = 'postgres' AND constraint_name ='primary'`
	query := `SELECT kcu.table_name, kcu.column_name, tco.constraint_name
	FROM information_schema.table_constraints tco
	JOIN information_schema.key_column_usage kcu 
		 ON kcu.constraint_name = tco.constraint_name
		 AND kcu.constraint_schema = tco.constraint_schema
		 AND kcu.constraint_name = tco.constraint_name
		 WHERE tco.constraint_type = 'PRIMARY KEY' ORDER BY kcu.table_schema, kcu.table_name`

	err := db.Select(&constraints, query)
	if err != nil {
		log.Fatal(err)
	}
	constraintMap := make(map[string][]string)
	for _, c := range constraints {
		constraintMap[c.TableName] = append(constraintMap[c.TableName], c.ColumnName)
	}
	return constraintMap
}

//Generate and return struct for one table
func GenerateModel(db *sqlx.DB, schema string, mi *ModelInfo) string {
	fields := GetTableFields(db, schema, *mi)
	members := []string{}
	fieldMap := map[string]string{}
	reverseFieldMap := map[string]string{}
	typeMap := map[string]string{}
	for _, f := range fields {
		// mi.HasId = mi.HasId || f.TableField == "id"
		member := fmt.Sprintf("%s %s `db:\"%s\" json:\",omitempty\"`", f.ModelField, f.Type, f.TableField)
		members = append(members, member)
		fieldMap[f.ModelField] = f.TableField
		reverseFieldMap[f.TableField] = f.ModelField
		typeMap[f.ModelField] = f.Type
	}
	st := fmt.Sprintf(`type %s struct {
		%s
	}
	`, mi.Model, strings.Join(members, "\n    "))

	fm, err := json.Marshal(fieldMap)
	if err != nil {
		log.Fatal(err)
	}
	rfm, err := json.Marshal(reverseFieldMap)
	if err != nil {
		log.Fatal(err)
	}
	tm, err := json.Marshal(typeMap)
	if err != nil {
		log.Fatal(err)
	}
	st = st + fmt.Sprintf(`func Get%sFieldMap() map[string]string {
		return map[string]string%s
		}

		`, mi.Model, fm)
	st = st + fmt.Sprintf(`func Get%sReverseFieldMap() map[string]string {
		return map[string]string%s
		}

		`, mi.Model, rfm)
	st = st + fmt.Sprintf(`func Get%sTypeMap() map[string]string {
		return map[string]string%s
		}

		`, mi.Model, tm)
	return st
}

func UpdatePkInfo(db *sqlx.DB, constraintsMap map[string][]string, mi *ModelInfo) {
	mi.HasId = len(constraintsMap[mi.Table]) == 1
	mi.Pk = constraintsMap[mi.Table]
}

//Generate for all tables
func GenerateModels(db *sqlx.DB, schema string, mis []*ModelInfo) {
	ms := []string{}
	constraintsMap := GetConstraintMap(db)
	for _, mi := range mis {
		UpdatePkInfo(db, constraintsMap, mi)
		fmt.Printf("Table: %s, HasId: %t\n", mi.Table, mi.HasId)
		m := GenerateModel(db, schema, mi)
		ms = append(ms, m)
	}
	content := []byte(`package models
	import (
		"time"

		"github.com/google/uuid"
		"github.com/senslabs/alpha/sens/datastore"
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
	t, err := template.New("fn.tpl").Funcs(template.FuncMap{
		"singular": func(s string) string {
			return strings.TrimSuffix(s, "s")
		},
	}).ParseFiles("templates/fn.tpl")

	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(fmt.Sprintf("generated/models/fn/%s.go", mi.Table))
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(f, types.Map{
		"Table": mi.Table,
		"Model": mi.Model,
		"HasId": mi.HasId,
		"Pk":    strings.Join(mi.Pk, ","),
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

// func GenerateDb(object string, model string) {
// 	t, err := template.ParseFiles("templates/db.tpl")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	f, err := os.Create(fmt.Sprintf("generated/api/db/%s.go", object))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = t.Execute(f, types.Map{"Model": model})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func GenerateApi(table string, model string, hasId bool) {
	t, err := template.ParseFiles("templates/rest.tpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(fmt.Sprintf("generated/api/%s.go", table))
	if err != nil {
		log.Fatal(err)
	}
	path := strings.ReplaceAll(table, "_", "-")
	err = t.Execute(f, types.Map{
		"Path":  path,
		"Model": model,
		"HasId": hasId,
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
	f, err := os.Create("generated/main/main.go")
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
		GenerateApi(mi.Table, mi.Model, mi.HasId)
		ms = append(ms, mi.Model)
	}
	GenerateMain(ms)
}

func main() {
	log.SetFlags(log.Lshortfile)
	schema := os.Args[1]
	os.Mkdir("generated/api", 0777)
	os.Mkdir("generated/main", 0777)
	os.Mkdir("generated/models", 0777)
	os.Mkdir("generated/models/fn", 0777)
	Generate(schema)
}
