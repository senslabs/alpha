package fn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
)

func InsertUserAuth(data []byte) (string, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return "", nil
	}
	var m models.UserAuth
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return "", nil
	}

	comma := ""
	fieldMap := models.GetUserAuthFieldMap()
	insert := bytes.NewBufferString("INSERT INTO user_auths(")
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
		logger.Error(err)
		return "", nil
	}
	var id string
	if err := stmt.Get(&id, m); err != nil {
		logger.Error(err)
		return "", nil
	} else {
		return id, nil
	}
}

func BatchInsertUserAuth(data []byte) ([]string, error) {
	var j []map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	var keys []string
	fieldMap := models.GetUserAuthFieldMap()
	insert := bytes.NewBufferString("INSERT INTO user_auths(")
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

	fmt.Fprint(insert, strings.TrimRight(ph.String(), ", ("), " returning id")

	fmt.Println(insert.String())

	db := models.GetConnection()
	_, err := db.Exec(insert.String(), values...)
	if err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.DB_ERROR, err)
	}
	return nil, nil
}

func UpdateUserAuth(id string, data []byte) error {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	var m models.UserAuth
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	fieldMap := models.GetUserAuthFieldMap()
	update := bytes.NewBufferString("UPDATE user_auths SET ")
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = :", f)
			comma = ", "
		}
	}
	fmt.Fprint(update, " WHERE id = :id")
	db := models.GetConnection()
	stmt, err := db.PrepareNamed(update.String())
	if err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	
	_, err = stmt.Exec(m)
	if err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	return nil
}

func SelectUserAuth(id string) (models.UserAuth, *errors.SensError) {
	db := models.GetConnection()
	m := models.UserAuth{}
	if err := db.Get(&m, "SELECT * FROM user_auths WHERE id = $1", id); err != nil {
		logger.Error(err)
		return m, errors.FromError(errors.DB_ERROR, err)
	}
	return m, nil
}

func FindUserAuth(or []string, and []string, span []string, limit string, column string, order string) ([]models.UserAuth, *errors.SensError) {
	ors := models.ParseOrParams(or)
	ands := models.ParseAndParams(and)
	spans := models.ParseSpanParams(span)

	fieldMap := models.GetUserAuthFieldMap()
	values := make(map[string]interface{})
	query := bytes.NewBufferString("SELECT * FROM user_auths WHERE ")
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
			fmt.Fprint(query, fmt.Sprintf("%s >= :from_%s AND %s <= :to_%s AND ", f, f, f, f))
			values["from_"+f] = s.From
			values["to_"+f] = s.To
		}
	}
	fmt.Fprint(query, "1 = 1)")
	if column != "" {
		if order == "" {
			order = "DESC"
		}
		fmt.Fprint(query, " ORDER BY ", column, " ", order)
	}
	fmt.Fprint(query, " LIMIT ", limit)
	
	m := []models.UserAuth{}
	db := models.GetConnection()
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
