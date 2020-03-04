package fn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/senslabs/modelgen/sens/models"
)

func InsertGroup(data []byte) (string, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		// logger.Error(err)
		log.Println(err)
		return "", nil
	}
	var m models.Group
	if err := json.Unmarshal(data, &m); err != nil {
		// logger.Error(err)
		log.Println(err)
		return "", nil
	}

	comma := ""
	fieldMap := models.GetGroupFieldMap()
	insert := bytes.NewBufferString("INSERT INTO groups(")
	values := bytes.NewBufferString("VALUES(")
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(insert, comma, f)
			fmt.Fprint(values, comma, ":", f)
			comma = ", "
		}
	}
	fmt.Fprint(insert, ") ")
	fmt.Fprint(insert, values, ") returning id")
	db := models.GetConnection()
	fmt.Println(insert.String())
	stmt, err := db.PrepareNamed(insert.String())
	if err != nil {
		// logger.Error(err)
		log.Println(err)
		return "", nil
	}
	var id string
	if err := stmt.Get(&id, m); err != nil {
		// logger.Error(err)
		log.Println(err)
		return "", nil
	} else {
		return id, nil
	}
}

func UpdateGroup(id string, data []byte) error {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		// logger.Error(err)
		log.Println(err)
		return err
	}
	var m models.Group
	if err := json.Unmarshal(data, &m); err != nil {
		// logger.Error(err)
		log.Println(err)
		return err
	}

	comma := ""
	fieldMap := models.GetGroupFieldMap()
	update := bytes.NewBufferString("UPDATE groups SET ")
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = :", f)
			comma = ", "
		}
	}
	db := models.GetConnection()
	fmt.Println(update)
	stmt, err := db.PrepareNamed(update.String())
	if err != nil {
		// logger.Error(err)
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(m)
	if err != nil {
		// logger.Error(err)
		log.Println(err)
		return err
	}
	return nil
}

func SelectGroup(id string) models.Group {
	db := models.GetConnection()
	m := models.Group{}
	err := db.Get(&m, "SELECT * FROM groups WHERE id = $1", id)
	if err != nil {
		log.Println(err)
	}
	return m
}

func FindGroup(or []string, and []string, span []string) []models.Group {
	ors := models.ParseOrParams(or)
	ands := models.ParseAndParams(and)
	spans := models.ParseSpanParams(span)

	fieldMap := models.GetGroupFieldMap()
	values := make(map[string]interface{})
	query := bytes.NewBufferString("SELECT * FROM groups WHERE ")
	for _, o := range ors {
		if f, ok := fieldMap[o.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s = :%s OR ", f, f))
			values[f] = o.Value
		}
	}
	fmt.Fprint(query, "(")
	for _, a := range ands {
		if f, ok := fieldMap[a.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s = :%s AND ", f, f))
			values[f] = a.Value
		}
	}
	for _, s := range spans {
		if f, ok := fieldMap[s.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s >= :from_%s AND %s &lt;= :to_%s AND ", f, f, f, f))
			values["from_"+f] = s.From
			values["to_"+f] = s.To
		}
	}
	fmt.Fprint(query, "1 = 1)")
	fmt.Println(query, values)
	m := []models.Group{}
	db := models.GetConnection()
	stmt, err := db.PrepareNamed(query.String())
	if err != nil {
		// logger.Error(err)
		log.Println(err)
		return m
	}
	if err := stmt.Select(&m, values); err != nil {
		// logger.Error(err)
		log.Println(err)
	}
	return m
}
