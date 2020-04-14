package fn

import (
	"bytes"
	"fmt"
	"runtime"
	"time"

	"github.com/lib/pq"
	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func InsertOrg(data []byte) string {
	j := types.UnmarshalMap(data)

	phi := 1
	comma := ""
	var values []interface{}
	fieldMap := models.GetOrgFieldMap()
	typeMap := models.GetOrgTypeMap()
	insert := bytes.NewBufferString("INSERT INTO orgs(")
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

	
	fmt.Fprint(insert, " returning org_id")
	

	db := datastore.GetConnection()

	logger.Debug(insert.String())

	stmt, err := db.Prepare(insert.String())
	errors.Pie(err)

	
	var id string
	r := stmt.QueryRow(values...)
	errors.Pie(r.Scan(&id))
	return id
	
}

func BatchInsertOrg(data []byte) {
	var j []map[string]interface{}
	types.Unmarshal(data, &j)

	comma := ""
	var keys []string
	fieldMap := models.GetOrgFieldMap()
	typeMap := models.GetOrgTypeMap()
	insert := bytes.NewBufferString("UPSERT INTO orgs(")
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
	errors.Pie(err)

	_, err = stmt.Exec(values...)
	errors.Pie(err)
}



func UpdateOrg(id string, data []byte) {
	var j map[string]interface{}
	types.Unmarshal(data, &j)

	phi := 1
	comma := ""
	var values []interface{}
	fieldMap := models.GetOrgFieldMap()
	update := bytes.NewBufferString("UPDATE orgs SET ")
	for k, v := range j {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(update, comma, f, " = $", phi)
			values = append(values, v)
			comma = ", "
			phi++
		}
	}
	values = append(values, id)
	fmt.Fprint(update, " WHERE alert_id = $", phi)

	logger.Debug(update.String())

	db := datastore.GetConnection()
	stmt, err := db.Prepare(update.String())
	errors.Pie(err)
	_, err = stmt.Exec(values...)
	errors.Pie(err)
}

func SelectOrg(id string) map[string]interface{} {
	db := datastore.GetConnection()

	stmt, err := db.Prepare("SELECT * FROM orgs WHERE alert_id = $1")
	errors.Pie(err)

	r, err := stmt.Query(id)
	errors.Pie(err)

	result := datastore.RowsToMap(r, models.GetOrgReverseFieldMap(), models.GetOrgTypeMap())
	if len(result) == 0 {
		return map[string]interface{}{}
	}
	return result[0]
}


func buildOrgWhereClause(query *bytes.Buffer, or []string, and []string, in string, span []string, values* []interface{}) {
	ors := datastore.ParseOrParams(or)
	ands := datastore.ParseAndParams(and)
	spans := datastore.ParseSpanParams(span)
	ins := datastore.ParseInParams(in)
	fieldMap := models.GetOrgFieldMap()

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

func FindOrg(or []string, and []string, in string, span []string, limit string, column string, order string) []map[string]interface{} {
	query := bytes.NewBufferString("SELECT * FROM orgs WHERE ")
	fieldMap := models.GetOrgFieldMap()
	var values []interface{}
	buildOrgWhereClause(query, or, and, in, span, &values)
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

	pc, file, line, ok := runtime.Caller(0)
	logger.Debug(time.Now().Unix(), "<BEFORE DB CONNECTION>", pc, file, line, ok)
	db := datastore.GetConnection()
	logger.Debug(time.Now().Unix(), "<AFTER DB CONNECTION>", pc, file, line, ok)
	stmt, err := db.Prepare(q)
	logger.Debug(time.Now().Unix(), "<AFTER PREPARE>", pc, file, line, ok)
	errors.Pie(err)

	r, err := stmt.Query(values...)
	logger.Debug(time.Now().Unix(), "<AFTER QUERY>", pc, file, line, ok)
	errors.Pie(err)

	result := datastore.RowsToMap(r, models.GetOrgReverseFieldMap(), models.GetOrgTypeMap())
	logger.Debug(time.Now().Unix(), "<RETURNING>", pc, file, line, ok)
	return result
}

func UpdateOrgWhere(or []string, and []string, in string, span []string, data []byte) {
	var values []interface{}
	j := types.UnmarshalMap(data)
	fieldMap := models.GetOrgFieldMap()
	typeMap := models.GetOrgTypeMap()
	update := bytes.NewBufferString("UPDATE orgs SET ")

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
	buildOrgWhereClause(update, or, and, in, span, &values)

	logger.Debug(update.String())
	logger.Debugf("Values: %#v", values)

	db := datastore.GetConnection()

	stmt, err := db.Prepare(update.String())
	errors.Pie(err)

	_, err = stmt.Query(values...)
	errors.Pie(err)
}