package fn

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/lib/pq"
	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func InsertSession(data []byte) string {
	j := types.UnmarshalMap(data)

	phi := 1
	comma := ""
	var values []interface{}
	fieldMap := models.GetSessionFieldMap()
	typeMap := models.GetSessionTypeMap()
	insert := bytes.NewBufferString("INSERT INTO sessions(")
	ph := bytes.NewBufferString("VALUES(")
	for k, v := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(insert, comma, f)
			fmt.Fprint(ph, comma, "$", phi)
			values = append(values, datastore.ConvertFieldValue(k, v, typeMap))
			comma = ", "
			phi++
		}
	}
	fmt.Fprint(insert, ") ")
	fmt.Fprint(insert, ph, ")")

	
	fmt.Fprint(insert, " returning session_id")
	

	db := datastore.GetConnection()

	logger.Debug(insert.String())

	stmt, err := db.Prepare(insert.String())
	defer stmt.Close()
	errors.Pie(err)

	
	var id string
	r := stmt.QueryRow(values...)
	errors.Pie(r.Scan(&id))
	return id
	
}

func BatchInsertSession(data []byte) {
	var j []map[string]interface{}
	types.Unmarshal(data, &j)

	comma := ""
	var keys []string
	fieldMap := models.GetSessionFieldMap()
	typeMap := models.GetSessionTypeMap()
	insert := bytes.NewBufferString("UPSERT INTO sessions(")
	for k, _ := range j[0] {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(insert, comma, f)
			keys = append(keys, k)
			comma = ", "
		}
	}

	phi := 1
	comma = ""
	var values []interface{}
	fmt.Fprint(insert, ") VALUES ")
	for _, kv := range j {
		fmt.Fprint(insert, comma, "(")
		comma = ""
		for _, k := range keys {
			values = append(values, datastore.ConvertFieldValue(k, kv[k], typeMap))
			fmt.Fprint(insert, comma, "$", phi)
			comma = ", "
			phi++
		}
		fmt.Fprint(insert, ")")
	}

	logger.Debug(insert.String())

	db := datastore.GetConnection()
	stmt, err := db.Prepare(insert.String())
	defer stmt.Close()
	errors.Pie(err)

	_, err = stmt.Exec(values...)
	errors.Pie(err)
}



func UpdateSession(id string, data []byte) {
	var j map[string]interface{}
	types.Unmarshal(data, &j)

	phi := 1
	comma := ""
	var values []interface{}
	fieldMap := models.GetSessionFieldMap()
	update := bytes.NewBufferString("UPDATE sessions SET ")
	for k, v := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = $", phi)
			values = append(values, v)
			comma = ", "
			phi++
		}
	}
	values = append(values, id)
	fmt.Fprint(update, " WHERE session_id = $", phi)

	logger.Debug(update.String())

	db := datastore.GetConnection()
	stmt, err := db.Prepare(update.String())
	defer stmt.Close()
	errors.Pie(err)
	_, err = stmt.Exec(values...)
	errors.Pie(err)
}

func SelectSession(id string) map[string]interface{} {
	db := datastore.GetConnection()

	stmt, err := db.Prepare("SELECT * FROM sessions WHERE session_id = $1")
	defer stmt.Close()
	errors.Pie(err)

	r, err := stmt.Query(id)
	errors.Pie(err)

	result := datastore.RowsToMap(r, models.GetSessionReverseFieldMap(), models.GetSessionTypeMap())
	if len(result) == 0 {
		return map[string]interface{}{}
	}
	return result[0]
}


func buildSessionWhereClause(query *bytes.Buffer, or []string, and []string, in string, span []string, values* []interface{}) {
	ors := datastore.ParseOrParams(or)
	ands := datastore.ParseAndParams(and)
	spans := datastore.ParseSpanParams(span)
	ins := datastore.ParseInParams(in)
	fieldMap := models.GetSessionFieldMap()

	phi := len(*values) + 1
	cond := ""
	fmt.Fprint(query, "(")
	for _, o := range ors {
		if f, ok := fieldMap[o.Column]; ok {
			fmt.Fprint(query, cond, f, " = $", phi)
			*values = append(*values, o.Value)
			cond = "OR "
			phi++
		}
	}

	if cond == "OR " {
		fmt.Fprint(query, ") AND (")
	}
	for _, a := range ands {
		if f, ok := fieldMap[a.Column]; ok {
			fmt.Fprint(query, f, " = $", phi, " AND ")
			*values = append(*values, a.Value)
			phi++
		}
	}

	if len(ins.Value) > 0 {
		if f, ok := fieldMap[ins.Column]; ok {
			fmt.Fprint(query, f, " = ANY($", phi, ") AND ")
			*values = append(*values, pq.Array(ins.Value))
			phi++
		}
	}

	for _, s := range spans {
		if f, ok := fieldMap[s.Column]; ok {
			if s.From != "" {
				fmt.Fprint(query, f, " >= $", phi, " AND ")
				*values = append(*values, s.From)
				phi++
			}
			if s.To != "" {
				fmt.Fprint(query, f, " <= $", phi, " AND ")
				*values = append(*values, s.To)
				phi++
			}
		}
	}
	fmt.Fprint(query, "1 = 1)")
}

func FindSession(or []string, and []string, in string, span []string, limit string, column string, order string) []map[string]interface{} {
	query := bytes.NewBufferString("SELECT * FROM sessions WHERE ")
	fieldMap := models.GetSessionFieldMap()
	var values []interface{}
	buildSessionWhereClause(query, or, and, in, span, &values)
	if column == "" {
		column = "created_at"
	}
	if f, ok := fieldMap[column]; ok {
		if order == "" {
			order = "DESC"
		}
		fmt.Fprint(query, " ORDER BY ", f, " ", order)
	}
	fmt.Fprint(query, " LIMIT ", limit)

	q := query.String()
	logger.Debug(q)

	seq := rand.Intn(99999)
	datastore.TRACE(seq, "1: <BEFORE DB CONNECTION>")
	db := datastore.GetConnection()
	datastore.TRACE(seq, "2: <AFTER DB CONNECTION>")
	stmt, err := db.Prepare(q)
	defer stmt.Close()
	datastore.TRACE(seq, "3: <AFTER PREPARED>")
	errors.Pie(err)

	r, err := stmt.Query(values...)
	datastore.TRACE(seq, "4: <AFTER QUERY>")
	errors.Pie(err)

	result := datastore.RowsToMap(r, models.GetSessionReverseFieldMap(), models.GetSessionTypeMap())
	datastore.TRACE(seq, "5: <RETURNING>")
	return result
}

func UpdateSessionWhere(or []string, and []string, in string, span []string, data []byte) {
	var values []interface{}
	j := types.UnmarshalMap(data)
	fieldMap := models.GetSessionFieldMap()
	typeMap := models.GetSessionTypeMap()
	update := bytes.NewBufferString("UPDATE sessions SET ")

	phi := 1
	comma := ""
	for k, v := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = $", phi)
			values = append(values, datastore.ConvertFieldValue(k, v, typeMap))
			comma = ", "
			phi++
		}
	}

	fmt.Fprint(update, " WHERE ")
	buildSessionWhereClause(update, or, and, in, span, &values)

	logger.Debug(update.String())
	logger.Debugf("Values: %#v", values)

	db := datastore.GetConnection()

	stmt, err := db.Prepare(update.String())
	defer stmt.Close()
	errors.Pie(err)

	_, err = stmt.Query(values...)
	errors.Pie(err)
}