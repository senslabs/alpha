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

func InsertSurveyAnswer(data []byte) (string, error) {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}
	var m models.SurveyAnswer
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return "", errors.FromError(errors.GO_ERROR, err)
	}

	logger.Debug(m)

	comma := ""
	fieldMap := models.GetSurveyAnswerFieldMap()
	insert := bytes.NewBufferString("INSERT INTO survey_answers(")
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
	
	fmt.Fprint(insert, " returning survey_answer_id")
	
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

func BatchInsertSurveyAnswer(data []byte) ([]string, error) {
	var j []map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}
	var m []*models.SurveyAnswer
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	var keys []string
	fieldMap := models.GetSurveyAnswerFieldMap()
	insert := bytes.NewBufferString("UPSERT INTO survey_answers(")
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


func UpdateSurveyAnswer(id string, data []byte) error {
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	var m models.SurveyAnswer
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}

	logger.Debug(m)

	comma := ""
	fieldMap := models.GetSurveyAnswerFieldMap()
	update := bytes.NewBufferString("UPDATE survey_answers SET ")
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = :", f)
			comma = ", "
		}
	}
	fmt.Fprint(update, " WHERE survey_answer_id = :survey_answer_id")

	logger.Debug(update.String())

	db := datastore.GetConnection()
	stmt, err := db.PrepareNamed(update.String())
	if err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	//<no value>
	m.SurveyAnswerId = &id
	_, err = stmt.Exec(m)
	if err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	return nil
}

func SelectSurveyAnswer(id string) (models.SurveyAnswer, *errors.SensError) {
	db := datastore.GetConnection()
	m := models.SurveyAnswer{}
	if err := db.Get(&m, "SELECT * FROM survey_answers WHERE survey_answer_id = $1", id); err != nil {
		logger.Error(err)
		return m, errors.FromError(errors.DB_ERROR, err)
	}
	return m, nil
}


func buildSurveyAnswerWhereClause(query *bytes.Buffer, or []string, and []string, in string, span []string, values map[string]interface{}) {
	ors := datastore.ParseOrParams(or)
	ands := datastore.ParseAndParams(and)
	spans := datastore.ParseSpanParams(span)
	ins := datastore.ParseInParams(in)
	fieldMap := models.GetSurveyAnswerFieldMap()

	cond := ""
	for _, o := range ors {
		if f, ok := fieldMap[o.Column]; ok {
			fmt.Fprint(query, cond, fmt.Sprintf("%s = :%s ", f, f))
			values[f] = getSurveyAnswerFieldValue(o.Column, o.Value)
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
			values[f] = getSurveyAnswerFieldValue(a.Column, a.Value)
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
			values["from_"+f] = getSurveyAnswerFieldValue(s.Column, s.From)
			values["to_"+f] = getSurveyAnswerFieldValue(s.Column, s.To)
		}
	}
	fmt.Fprint(query, "1 = 1)")
}

func getSurveyAnswerFieldValue(c string, v interface{}) interface{} {
	typeMap := models.GetSurveyAnswerTypeMap()
	if typeMap[c] == "*datastore.RawMessage" {
		v, _ = json.Marshal(v)
	}
	return v
}

func findSurveyAnswerIn(query string, values map[string]interface{}) ([]models.SurveyAnswer, *errors.SensError) {
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
		m := []models.SurveyAnswer{}
		if err := db.Select(&m, q, a...); err != nil {
			logger.Error(err)
			return m, errors.New(errors.DB_ERROR, err.Error())
		}
		return m, nil
	}
}

func FindSurveyAnswer(or []string, and []string, in string, span []string, limit string, column string, order string) ([]models.SurveyAnswer, *errors.SensError) {
	from := time.Now().Unix()
	query := bytes.NewBufferString("SELECT * FROM survey_answers WHERE ")
	fieldMap := models.GetSurveyAnswerFieldMap()
	values := make(map[string]interface{})
	buildSurveyAnswerWhereClause(query, or, and, in, span, values)
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

		var m []models.SurveyAnswer
		if err := stmt.Select(&m, values); err != nil {
			logger.Error(err)
			return nil, errors.New(errors.DB_ERROR, err.Error())
		} else {
			to := time.Now().Unix()
			logger.Debugf("Returning FIND after %d seconds: RESULT => %#v", (to - from), m)
			return m, nil
		}
	} else {
		logger.Debug("Before find In")
		m, err := findSurveyAnswerIn(q, values)
		logger.Debug("After find In")
		to := time.Now().Unix()
		logger.Debugf("Returning IN after %d seconds: RESULT => %#v", (to - from), m)
		return m, err
	}
}

func UpdateSurveyAnswerWhere(or []string, and []string, in string, span []string, data []byte) *errors.SensError {
	fieldMap := models.GetSurveyAnswerFieldMap()
	values := make(map[string]interface{})
	update := bytes.NewBufferString("UPDATE survey_answers SET ")

	//SET FIELD VALUES
	var j map[string]interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	var m models.SurveyAnswer
	if err := json.Unmarshal(data, &m); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}

	comma := ""
	for k, _ := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = :set_", f)
			values["set_"+f] = getSurveyAnswerFieldValue(k, j[k])
			comma = ", "
		}
	}
	//SET ENDS

	fmt.Fprint(update, " WHERE ")
	buildSurveyAnswerWhereClause(update, or, and, in, span, values)

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
