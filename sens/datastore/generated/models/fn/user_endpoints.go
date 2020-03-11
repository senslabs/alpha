package fn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
)

func InsertUserEndpoint(data []byte) (string, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return "", nil
	}
	var m models.UserEndpoint
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return "", nil
	}

	comma := ""
	fieldMap := models.GetUserEndpointFieldMap()
	insert := bytes.NewBufferString("INSERT INTO user_endpoints(")
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

func UpdateUserEndpoint(id string, data []byte) error {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	var m models.UserEndpoint
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	fieldMap := models.GetUserEndpointFieldMap()
	update := bytes.NewBufferString("UPDATE user_endpoints SET ")
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

func SelectUserEndpoint(id string) (models.UserEndpoint, *errors.SensError) {
	db := models.GetConnection()
	m := models.UserEndpoint{}
	if err := db.Get(&m, "SELECT * FROM auths WHERE id = $1", id); err != nil {
		logger.Error(err)
		return m, errors.FromError(errors.DB_ERROR, err)
	}
	return m, nil
}

func FindUserEndpoint(or []string, and []string, span []string, limit string, column string, order string) ([]models.UserEndpoint, *errors.SensError) {
	ors := models.ParseOrParams(or)
	ands := models.ParseAndParams(and)
	spans := models.ParseSpanParams(span)

	fieldMap := models.GetUserEndpointFieldMap()
	values := make(map[string]interface{})
	query := bytes.NewBufferString("SELECT * FROM user_endpoints WHERE ")
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
	
	m := []models.UserEndpoint{}
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
