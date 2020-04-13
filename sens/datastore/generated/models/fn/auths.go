package fn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/jmoiron/sqlx"
)

func InsertAuth(data []byte) (string, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}
	var m models.Auth
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}

	logger.Debug(m)

	comma := ""
	fieldMap := models.GetAuthFieldMap()
	insert := bytes.NewBufferString("INSERT INTO auths(")
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
	
	fmt.Fprint(insert, " returning auth_id")
	
	db := datastore.GetConnection()

	logger.Debug(insert.String())
	logger.Debugf("%#v", m)

	stmt, err := db.PrepareNamed(insert.String())
	if err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.DB_ERROR, err)
	}
	
	var id string
	if err := stmt.Get(&id, m); err != nil {
		logger.Errorf("Received error %s while inserting values\n\t %#v", err, values)
		return "", errors.FromError(errors.DB_ERROR, err)
	} else {
		return id, nil
	}
	
}

func BatchInsertAuth(data []byte) ([]string, error) {
	var j []map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}
	var m []*models.Auth
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	var keys []string
	fieldMap := models.GetAuthFieldMap()
	insert := bytes.NewBufferString("UPSERT INTO auths(")
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


func UpdateAuth(id string, data []byte) error {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	var m models.Auth
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}

	logger.Debug(m)

	comma := ""
	fieldMap := models.GetAuthFieldMap()
	update := bytes.NewBufferString("UPDATE auths SET ")
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = :", f)
			comma = ", "
		}
	}
	fmt.Fprint(update, " WHERE auth_id = :auth_id")

	logger.Debug(update.String())

	db := datastore.GetConnection()
	stmt, err := db.PrepareNamed(update.String())
	if err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	//<no value>
	m.AuthId = &id
	_, err = stmt.Exec(m)
	if err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	return nil
}

func SelectAuth(id string) (models.Auth, *errors.SensError) {
	db := datastore.GetConnection()
	m := models.Auth{}
	if err := db.Get(&m, "SELECT * FROM auths WHERE auth_id = $1", id); err != nil {
		logger.Error(err)
		return m, errors.FromError(errors.DB_ERROR, err)
	}
	return m, nil
}


func buildAuthWhereClause(query *bytes.Buffer, or []string, and []string, in string, span []string, values map[string]interface{}) {
	ors := datastore.ParseOrParams(or)
	ands := datastore.ParseAndParams(and)
	spans := datastore.ParseSpanParams(span)
	ins := datastore.ParseInParams(in)
	fieldMap := models.GetAuthFieldMap()

	cond := ""
	for _, o := range ors {
		if f, ok := fieldMap[o.Column]; ok {
			fmt.Fprint(query, cond, fmt.Sprintf("%s = :%s ", f, f))
			values[f] = getAuthFieldValue(o.Column, o.Value)
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
			values[f] = getAuthFieldValue(a.Column, a.Value)
		}
	}

	if len(ins.Value) > 0 {
		if f, ok := fieldMap[ins.Column]; ok {
			fmt.Fprint(query, f, " in (:", f, ") AND ")
			values[f] = ins.Value
		}
	}

	for _, s := range spans {
		if f, ok := fieldMap[s.Column]; ok {
			fmt.Fprint(query, fmt.Sprintf("%s >= :from_%s AND %s <= :to_%s AND ", f, f, f, f))
			values["from_"+f] = getAuthFieldValue(s.Column, s.From)
			values["to_"+f] = getAuthFieldValue(s.Column, s.To)
		}
	}
	fmt.Fprint(query, "1 = 1)")
}

func getAuthFieldValue(c string, v interface{}) interface{} {
	typeMap := models.GetAuthTypeMap()
	if typeMap[c] == "*datastore.RawMessage" {
		v, _ = json.Marshal(v)
	}
	return v
}

func findAuthIn(query string, values map[string]interface{}) ([]models.Auth, *errors.SensError) {
	if q, a, err := sqlx.Named(query, values); err != nil {
		logger.Error(err.Error())
		return nil, errors.New(errors.DB_ERROR, err.Error())
	} else if q, a, err := sqlx.In(q, a...); err != nil {
		logger.Error(err.Error())
		return nil, errors.New(errors.DB_ERROR, err.Error())
	} else {
		db := datastore.GetConnection()
		q = db.Rebind(q)
		logger.Debug(q)
		logger.Debugf("Values: %s", a)
		m := []models.Auth{}
		if err := db.Select(&m, q, a...); err != nil {
			logger.Error(err)
			return m, errors.New(errors.DB_ERROR, err.Error())
		}
		return m, nil
	}
}

func FindAuth(or []string, and []string, in string, span []string, limit string, column string, order string) ([]models.Auth, *errors.SensError) {
	from := time.Now().Unix()
	seq := from % 10000
	query := bytes.NewBufferString("SELECT * FROM auths WHERE ")
	fieldMap := models.GetAuthFieldMap()
	values := make(map[string]interface{})
	buildAuthWhereClause(query, or, and, in, span, values)
	if column != "" {
		if f, ok := fieldMap[column]; ok {
			if order == "" {
				order = "DESC"
			}
			fmt.Fprint(query, " ORDER BY ", f, " ", order)
		}
	}
	fmt.Fprint(query, " LIMIT ", limit)

	q := query.String()
	logger.Debug(q)
	if strings.TrimSpace(in) == "" {
		logger.Debug("No in clause present. Using prepared and not sqlIn")
		db := datastore.GetConnection()
		stmt, err := db.PrepareNamed(q)
		if err != nil {
			logger.Error(err.Error())
			return nil, errors.New(errors.DB_ERROR, err.Error())
		}

		var m []models.Auth
		if err := stmt.Select(&m, values); err != nil {
			logger.Error(err)
			return nil, errors.New(errors.DB_ERROR, err.Error())
		} else {
			to := time.Now().Unix()
			logger.Debugf("%d: Returning FIND after %d seconds: RESULT => %#v", seq, (to - from), m)
			return m, nil
		}
	} else {
		logger.Debug(seq, ": Before find In")
		m, err := findAuthIn(q, values)
		logger.Debug(seq, " :After find In")
		to := time.Now().Unix()
		logger.Debugf("%d: Returning IN after %d seconds: RESULT => %#v", seq, (to - from), m)
		return m, err
	}
}

func UpdateAuthWhere(or []string, and []string, in string, span []string, data []byte) *errors.SensError {
	fieldMap := models.GetAuthFieldMap()
	values := make(map[string]interface{})
	update := bytes.NewBufferString("UPDATE auths SET ")

	//SET FIELD VALUES
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	var m models.Auth
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = :set_", f)
			values["set_"+f] = getAuthFieldValue(k, j[k])
			comma = ", "
		}
	}
	//SET ENDS

	fmt.Fprint(update, " WHERE ")
	buildAuthWhereClause(update, or, and, in, span, values)

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
