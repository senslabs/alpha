package fn

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
)

func InsertUserAlertView(data []byte) (string, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}
	var m models.UserAlertView
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}

	logger.Debug(m)

	comma := ""
	fieldMap := models.GetUserAlertViewFieldMap()
	insert := bytes.NewBufferString("INSERT INTO user_alert_views(")
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
	logger.Debugf("%#v", m)

	stmt, err := db.PrepareNamed(insert.String())
	if err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.DB_ERROR, err)
	}
	
	if _, err := stmt.Exec(m); err != nil {
		logger.Errorf("Received error %s while inserting values\n\t %#v", err, values)
		return "", errors.FromError(errors.DB_ERROR, err)
	} else {
		return "", nil
	}
	
}

func BatchInsertUserAlertView(data []byte) ([]string, error) {
	var j []map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}
	var m []*models.SessionRecord
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	var keys []string
	fieldMap := models.GetUserAlertViewFieldMap()
	insert := bytes.NewBufferString("UPSERT INTO user_alert_views(")
	ph := bytes.NewBufferString("(")
	for k, _ := range j[0] {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(insert, comma, f)
			fmt.Fprint(ph, comma, ":", f)
			keys = append(keys, k)
			comma = ", "
		}
	}
	fmt.Fprint(ph, ")")
	fmt.Fprint(insert, ") VALUES ")

	fmt.Fprint(insert, ph.String())

	logger.Debug(insert.String())

	db := datastore.GetConnection()
	if _, err := db.NamedExec(insert.String(), m); err != nil {
		logger.Errorf("Received error %s while inserting values\n\t %#v", err, m)
		return nil, errors.FromError(errors.DB_ERROR, err)
	}
	return nil, nil
}



func buildUserAlertViewWhereClause(query *bytes.Buffer, or []string, and []string, span []string, values map[string]interface{}) {
	ors := datastore.ParseOrParams(or)
	ands := datastore.ParseAndParams(and)
	spans := datastore.ParseSpanParams(span)
	fieldMap := models.GetUserAlertViewFieldMap()

	cond := ""
	for _, o := range ors {
		if f, ok := fieldMap[o.Column]; ok {
			fmt.Fprint(query, cond, fmt.Sprintf("%s = :%s ", f, f))
			values[f] = getUserAlertViewFieldValue(o.Column, o.Value)
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
			values[f] = getUserAlertViewFieldValue(a.Column, a.Value)
		}
	}
	for _, s := range spans {
		if f, ok := fieldMap[s.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s >= :from_%s AND %s <= :to_%s AND ", f, f, f, f))
			values["from_"+f] = getUserAlertViewFieldValue(s.Column, s.From)
			values["to_"+f] = getUserAlertViewFieldValue(s.Column, s.To)
		}
	}
	fmt.Fprint(query, "1 = 1)")
}

func getUserAlertViewFieldValue(c string, v interface{}) interface{} {
	// typeMap := models.GetAuthTypeMap()
	return v
}

func FindUserAlertView(or []string, and []string, span []string, limit string, column string, order string) ([]models.UserAlertView, *errors.SensError) {
	query := bytes.NewBufferString("SELECT * FROM user_alert_views WHERE ")
	fieldMap := models.GetUserAlertViewFieldMap()
	values := make(map[string]interface{})
	buildUserAlertViewWhereClause(query, or, and, span, values)
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

	m := []models.UserAlertView{}
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

func UpdateUserAlertViewWhere(or []string, and []string, span []string, data []byte) *errors.SensError {
	fieldMap := models.GetUserAlertViewFieldMap()
	values := make(map[string]interface{})
	update := bytes.NewBufferString("UPDATE user_alert_views SET ")

	//SET FIELD VALUES
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	var m models.UserAlertView
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = :set_", f)
			values["set_"+f] = j[k]
			comma = ", "
		}
	}
	//SET ENDS

	fmt.Fprint(update, " WHERE ")
	buildUserAlertViewWhereClause(update, or, and, span, values)

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
