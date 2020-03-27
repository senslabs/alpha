package fn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
)

func InsertOpEndpointCategorie(data []byte) (string, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}
	var m models.OpEndpointCategorie
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}

	logger.Debug(m)

	comma := ""
	fieldMap := models.GetOpEndpointCategorieFieldMap()
	insert := bytes.NewBufferString("INSERT INTO op_endpoint_categories(")
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
	
	db := datastore.GetConnection()

	logger.Debug(insert.String())
	
	stmt, err := db.PrepareNamed(insert.String())
	if err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.DB_ERROR, err)
	}
	
	if _, err := stmt.Exec(m); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.DB_ERROR, err)
	} else {
		return "", nil
	}
	
}

func BatchInsertOpEndpointCategorie(data []byte) ([]string, error) {
	var j []map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	var keys []string
	fieldMap := models.GetOpEndpointCategorieFieldMap()
	insert := bytes.NewBufferString("INSERT INTO op_endpoint_categories(")
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



func buildOpEndpointCategorieWhereClause(query *bytes.Buffer, or []string, and []string, span []string, values map[string]interface{}) {
	ors := datastore.ParseOrParams(or)
	ands := datastore.ParseAndParams(and)
	spans := datastore.ParseSpanParams(span)
	fieldMap := models.GetOpEndpointCategorieFieldMap()

	cond := ""
	for _, o := range ors {
		if f, ok := fieldMap[o.Column]; ok {
			fmt.Fprint(query, cond, fmt.Sprintf("%s = :%s ", f, f))
			values[f] = getOpEndpointCategorieFieldValue(o.Column, o.Value)
			cond = "OR "
		}
	}

	if cond == "OR " {
		fmt.Fprint(query, "AND ")
	}
	fmt.Fprint(query, "(")
	for _, a := range ands {
		if f, ok := fieldMap[a.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s = :%s AND ", f, f))
			values[f] = getOpEndpointCategorieFieldValue(a.Column, a.Value)
		}
	}
	for _, s := range spans {
		if f, ok := fieldMap[s.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s >= :from_%s AND %s <= :to_%s AND ", f, f, f, f))
			values["from_"+f] = getOpEndpointCategorieFieldValue(s.Column, s.From)
			values["to_"+f] = getOpEndpointCategorieFieldValue(s.Column, s.To)
		}
	}
	fmt.Fprint(query, "1 = 1)")
}

func getOpEndpointCategorieFieldValue(c string, v interface{}) interface{} {
	// typeMap := models.GetAuthTypeMap()
	return v
}

func FindOpEndpointCategorie(or []string, and []string, span []string, limit string, column string, order string) ([]models.OpEndpointCategorie, *errors.SensError) {
	query := bytes.NewBufferString("SELECT * FROM op_endpoint_categories WHERE ")
	fieldMap := models.GetOpEndpointCategorieFieldMap()
	values := make(map[string]interface{})
	buildOpEndpointCategorieWhereClause(query, or, and, span, values)
	if column != "" {
		if f, ok := fieldMap[column]; ok {
			if order == "" {
				order = "DESC"
			}
			fmt.Fprint(query, " ORDER BY ", f, " ", order)
		}
	}
	fmt.Fprint(query, " LIMIT ", limit)

	logger.Debug(query.String())
	logger.Debugf("Values: %#v", values)

	m := []models.OpEndpointCategorie{}
	db := datastore.GetConnection()
	if stmt, err := db.PrepareNamed(query.String()); err != nil {
		logger.Error(err.Error())
		return m, errors.New(errors.DB_ERROR, err.Error())
	} else if err := stmt.Select(&m, values); err != nil {
		logger.Error(err)
		return m, errors.New(errors.DB_ERROR, err.Error())
	}
	return m, nil
}

func UpdateOpEndpointCategorieWhere(or []string, and []string, span []string, data []byte) *errors.SensError {
	fieldMap := models.GetOpEndpointCategorieFieldMap()
	values := make(map[string]interface{})
	update := bytes.NewBufferString("UPDATE op_endpoint_categories SET ")

	//SET FIELD VALUES
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	var m models.OpEndpointCategorie
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = :set_", k)
			values["set_"+f] = j[f]
			comma = ", "
		}
	}
	//SET ENDS

	fmt.Fprint(update, " WHERE ")
	buildOpEndpointCategorieWhereClause(update, or, and, span, values)

	logger.Debug(update.String())
	logger.Debugf("Values: %#v", values)

	db := datastore.GetConnection()
	stmt, err := db.PrepareNamed(update.String())
	if err != nil {
		logger.Error(err.Error())
		return errors.New(errors.DB_ERROR, err.Error())
	}

	if _, err := stmt.Exec(values); err != nil {
		logger.Error(err)
		return errors.New(errors.DB_ERROR, err.Error())
	}
	return nil
}
