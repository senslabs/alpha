package fn

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func InsertOpUser(data []byte) string {
	j := types.UnmarshalMap(data)
	if len (j) == 0 {
		errors.Pie(errors.New(0, "NO DATA"))
	}

	phi := 1
	comma := ""
	var values []interface{}
	fieldMap := models.GetOpUserFieldMap()
	typeMap := models.GetOpUserTypeMap()
	insert := bytes.NewBufferString("INSERT INTO op_users(")
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

	

	db := datastore.GetConnection()

	logger.Debug(insert.String())

	stmt, err := db.Prepare(insert.String())
	defer stmt.Close()
	errors.Pie(err)

	
	_, err = stmt.Exec(values...)
	errors.Pie(err)
	return ""
	
}

func BatchInsertOpUser(data []byte) {
	var j []map[string]interface{}
	types.Unmarshal(data, &j)
	if len(j) == 0 {
		return
	}

	comma := ""
	var keys []string
	var fields []string
	fieldMap := models.GetOpUserFieldMap()
	typeMap := models.GetOpUserTypeMap()
	insert := bytes.NewBufferString("INSERT INTO op_users(")
	for k, _ := range j[0] {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(insert, comma, f)
			keys = append(keys, k)
			fields = append(fields, f)
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

	fmt.Fprint(insert, " ON CONFLICT(op_id,user_id) DO UPDATE SET (", strings.Join(fields, ", "), ") = (EXCLUDED.", strings.Join(fields, ", EXCLUDED."), ")")


	logger.Debug(insert.String())

	db := datastore.GetConnection()
	stmt, err := db.Prepare(insert.String())
	defer stmt.Close()
	errors.Pie(err)

	_, err = stmt.Exec(values...)
	errors.Pie(err)
}

func BatchUpsertOpUser(data []byte) {
	var j []map[string]interface{}
	types.Unmarshal(data, &j)
	if len(j) == 0 {
		return
	}

	comma := ""
	var keys []string
	var fields []string
	fieldMap := models.GetOpUserFieldMap()
	typeMap := models.GetOpUserTypeMap()
	insert := bytes.NewBufferString("UPSERT INTO op_users(")
	for k, _ := range j[0] {
		if f, ok := fieldMap[k]; ok {
			fmt.Fprint(insert, comma, f)
			keys = append(keys, k)
			fields = append(fields, f)
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

	//fmt.Fprint(insert, " ON CONFLICT(op_id,user_id) DO UPDATE SET (", strings.Join(fields, ", "), ") = (EXCLUDED.", strings.Join(fields, ", EXCLUDED."), ")")

	logger.Debug(insert.String())

	db := datastore.GetConnection()
	stmt, err := db.Prepare(insert.String())
	defer stmt.Close()
	errors.Pie(err)

	_, err = stmt.Exec(values...)
	errors.Pie(err)
}




func buildOpUserWhereClause(query *bytes.Buffer, or []string, and []string, in string, span []string, values* []interface{}) {
	ors := datastore.ParseOrParams(or)
	ands := datastore.ParseAndParams(and)
	spans := datastore.ParseSpanParams(span)
	ins := datastore.ParseInParams(in)
	fieldMap := models.GetOpUserFieldMap()

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

func FindOpUser(or []string, and []string, in string, span []string, limit string, column string, order string) []map[string]interface{} {
	query := bytes.NewBufferString("SELECT * FROM op_users WHERE ")
	fieldMap := models.GetOpUserFieldMap()
	var values []interface{}
	buildOpUserWhereClause(query, or, and, in, span, &values)
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

	db := datastore.GetConnection()
	stmt, err := db.Prepare(q)
	defer stmt.Close()
	errors.Pie(err)

	r, err := stmt.Query(values...)
	errors.Pie(err)

	result := datastore.RowsToMap(r, models.GetOpUserReverseFieldMap(), models.GetOpUserTypeMap())
	return result
}

func UpdateOpUserWhere(or []string, and []string, in string, span []string, data []byte) {
	var values []interface{}
	j := types.UnmarshalMap(data)
	fieldMap := models.GetOpUserFieldMap()
	typeMap := models.GetOpUserTypeMap()
	update := bytes.NewBufferString("UPDATE op_users SET ")

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
	buildOpUserWhereClause(update, or, and, in, span, &values)

	logger.Debug(update.String())
	logger.Debugf("Values: %#v", values)

	db := datastore.GetConnection()

	stmt, err := db.Prepare(update.String())
	defer stmt.Close()
	errors.Pie(err)

	_, err = stmt.Query(values...)
	errors.Pie(err)
}

func DeleteOpUser(or []string, and []string, in string, span []string) int64 {
	query := bytes.NewBufferString("DELETE FROM op_users WHERE ")
	var values []interface{}
	buildOpUserWhereClause(query, or, and, in, span, &values)
	q := query.String()
	logger.Debug(q)

	db := datastore.GetConnection()
	stmt, err := db.Prepare(q)
	defer stmt.Close()
	errors.Pie(err)

	r, err := stmt.Exec(values...)
	errors.Pie(err)

	n, err := r.RowsAffected()
	errors.Pie(err)
	return n
}