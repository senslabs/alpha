package fn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
)

func Insert{{.Model}}(data []byte) (string, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}
	var m models.{{.Model}}
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}

	logger.Debug(m)

	comma := ""
	fieldMap := models.Get{{.Model}}FieldMap()
	insert := bytes.NewBufferString("INSERT INTO {{.Table}}(")
	values := bytes.NewBufferString("VALUES(")
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(insert, comma, f)
			fmt.Fprint(values, comma, ":", f)
			comma = ", "
		}
	}
	fmt.Fprint(insert, ") ")
	fmt.Fprint(insert, values, ")")
	{{if .HasId}}
	fmt.Fprint(insert, " returning id")
	{{end}}
	db := datastore.GetConnection()

	logger.Debug(insert.String())
	
	stmt, err := db.PrepareNamed(insert.String())
	if err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.DB_ERROR, err)
	}
	{{if .HasId}}
	var id string
	if err := stmt.Get(&id, m); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.DB_ERROR, err)
	} else {
		return id, nil
	}
	{{else}}
	if _, err := stmt.Exec(m); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.DB_ERROR, err)
	} else {
		return "", nil
	}
	{{end}}
}

func BatchInsert{{.Model}}(data []byte) ([]string, error) {
	var j []map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	var keys []string
	fieldMap := models.Get{{.Model}}FieldMap()
	insert := bytes.NewBufferString("INSERT INTO {{.Table}}(")
	for k, _ := range j[0] {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(insert, comma, f)
			keys = append(keys, k)
			comma = ", "
		}
	}

	ph := bytes.NewBufferString(") VALUES (")
	phidx := 1
	var values []interface{}
	for _, v := range j {
		comma = ""
		for _, k := range keys {
			fmt.Fprint(ph, comma, "$", phidx)
			phidx++
			values = append(values, v[k])
			comma = ", "
		}
		fmt.Fprint(ph, "), (")
	}

	fmt.Fprint(insert, strings.TrimRight(ph.String(), ", ("))

	logger.Debug(insert.String())

	db := datastore.GetConnection()
	_, err := db.Exec(insert.String(), values...)
	if err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.DB_ERROR, err)
	}
	return nil, nil
}

{{if .HasId}}
func Update{{.Model}}(id string, data []byte) error {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	var m models.{{.Model}}
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}

	logger.Debug(m)

	comma := ""
	fieldMap := models.Get{{.Model}}FieldMap()
	update := bytes.NewBufferString("UPDATE {{.Table}} SET ")
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = :", f)
			comma = ", "
		}
	}
	fmt.Fprint(update, " WHERE id = :id")

	logger.Debug(update.String())

	db := datastore.GetConnection()
	stmt, err := db.PrepareNamed(update.String())
	if err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	//{{.Id}}
	m.Id = id
	_, err = stmt.Exec(m)
	if err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	return nil
}

func Select{{.Model}}(id string) (models.{{.Model}}, *errors.SensError) {
	db := datastore.GetConnection()
	m := models.{{.Model}}{}
	if err := db.Get(&m, "SELECT * FROM {{.Table}} WHERE id = $1", id); err != nil {
		logger.Error(err)
		return m, errors.FromError(errors.DB_ERROR, err)
	}
	return m, nil
}
{{end}}

func get{{.Model}}FieldValue(c string, v interface{}) interface{} {
	typeMap := models.Get{{.Model}}TypeMap()
	if typeMap[c] == "datastore.NullTime" || typeMap[c] == "TIMESTAMP" {
		if val, err := strconv.ParseInt(v.(string), 10, 64); err != nil {
			logger.Error(err)
		} else {
			return time.Unix(val, 0)
		}
	}
	return v
}

func Find{{.Model}}(or []string, and []string, span []string, limit string, column string, order string) ([]models.{{.Model}}, *errors.SensError) {
	ors := datastore.ParseOrParams(or)
	ands := datastore.ParseAndParams(and)
	spans := datastore.ParseSpanParams(span)

	fieldMap := models.Get{{.Model}}FieldMap()
	values := make(map[string]interface{})
	query := bytes.NewBufferString("SELECT * FROM {{.Table}} WHERE ")
	for _, o := range ors {
		if f, ok := fieldMap[o.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s = :%s OR ", f, f))
			values[f] = get{{.Model}}FieldValue(o.Column, o.Value)
		}
	}
	fmt.Fprint(query, "(")
	for _, a := range ands {
		if f, ok := fieldMap[a.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s = :%s AND ", f, f))
			values[f] = get{{.Model}}FieldValue(a.Column, a.Value)
		}
	}
	for _, s := range spans {
		if f, ok := fieldMap[s.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s >= :from_%s AND %s <= :to_%s AND ", f, f, f, f))
			values["from_"+f] = get{{.Model}}FieldValue(s.Column, s.From)
			values["to_"+f] = get{{.Model}}FieldValue(s.Column, s.To)
		}
	}
	fmt.Fprint(query, "1 = 1)")
	if column != "" {
		if f, ok := fieldMap[column]; ok {
			if order == "" {
				order = "DESC"
			}
			fmt.Fprint(query, " ORDER BY :", f, " ", order)
			values[column] = f
		}
	}
	fmt.Fprint(query, " LIMIT ", limit)

	logger.Debug(query.String())
	
	m := []models.{{.Model}}{}
	db := datastore.GetConnection()
	if stmt, err := db.PrepareNamed(query.String()); err != nil {
		logger.Error(err.Error())
		log.Printf("%v", err)
		return m, errors.New(errors.DB_ERROR, err.Error())
	} else if err := stmt.Select(&m, values); err != nil {
		logger.Error(err)
		return m, errors.New(errors.DB_ERROR, err.Error())
	}
	return m, nil
}
