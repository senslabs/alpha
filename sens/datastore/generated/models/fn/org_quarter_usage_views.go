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

func InsertOrgQuarterUsageView(data []byte) string {
	j := types.UnmarshalMap(data)
	if len (j) == 0 {
		errors.Pie(errors.New(0, "NO DATA"))
	}

	phi := 1
	comma := ""
	var values []interface{}
	fieldMap := models.GetOrgQuarterUsageViewFieldMap()
	typeMap := models.GetOrgQuarterUsageViewTypeMap()
	insert := bytes.NewBufferString("INSERT INTO org_quarter_usage_views(")
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

func BatchInsertOrgQuarterUsageView(data []byte) {
	var j []map[string]interface{}
	types.Unmarshal(data, &j)
	if len(j) == 0 {
		return
	}

	comma := ""
	var keys []string
	var fields []string
	fieldMap := models.GetOrgQuarterUsageViewFieldMap()
	typeMap := models.GetOrgQuarterUsageViewTypeMap()
	insert := bytes.NewBufferString("INSERT INTO org_quarter_usage_views(")
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

	fmt.Fprint(insert, " ON CONFLICT() DO UPDATE SET (", strings.Join(fields, ", "), ") = (EXCLUDED.", strings.Join(fields, ", EXCLUDED."), ")")


	logger.Debug(insert.String())

	db := datastore.GetConnection()
	stmt, err := db.Prepare(insert.String())
	defer stmt.Close()
	errors.Pie(err)

	_, err = stmt.Exec(values...)
	errors.Pie(err)
}

func BatchUpsertOrgQuarterUsageView(data []byte) {
	var j []map[string]interface{}
	types.Unmarshal(data, &j)
	if len(j) == 0 {
		return
	}

	comma := ""
	var keys []string
	var fields []string
	fieldMap := models.GetOrgQuarterUsageViewFieldMap()
	typeMap := models.GetOrgQuarterUsageViewTypeMap()
	insert := bytes.NewBufferString("UPSERT INTO org_quarter_usage_views(")
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

	//fmt.Fprint(insert, " ON CONFLICT() DO UPDATE SET (", strings.Join(fields, ", "), ") = (EXCLUDED.", strings.Join(fields, ", EXCLUDED."), ")")

	logger.Debug(insert.String())

	db := datastore.GetConnection()
	stmt, err := db.Prepare(insert.String())
	defer stmt.Close()
	errors.Pie(err)

	_, err = stmt.Exec(values...)
	errors.Pie(err)
}




func buildOrgQuarterUsageViewWhereClause(query *bytes.Buffer, or []string, and []string, in string, span []string, values* []interface{}) {
	ors := datastore.ParseOrParams(or)
	ands := datastore.ParseAndParams(and)
	spans := datastore.ParseSpanParams(span)
	ins := datastore.ParseInParams(in)
	fieldMap := models.GetOrgQuarterUsageViewFieldMap()

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

func FindOrgQuarterUsageView(or []string, and []string, in string, span []string, limit string, column string, order string) []map[string]interface{} {
	query := bytes.NewBufferString("SELECT * FROM org_quarter_usage_views WHERE ")
	fieldMap := models.GetOrgQuarterUsageViewFieldMap()
	var values []interface{}
	buildOrgQuarterUsageViewWhereClause(query, or, and, in, span, &values)
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

	result := datastore.RowsToMap(r, models.GetOrgQuarterUsageViewReverseFieldMap(), models.GetOrgQuarterUsageViewTypeMap())
	return result
}

func UpdateOrgQuarterUsageViewWhere(or []string, and []string, in string, span []string, data []byte) {
	var values []interface{}
	j := types.UnmarshalMap(data)
	fieldMap := models.GetOrgQuarterUsageViewFieldMap()
	typeMap := models.GetOrgQuarterUsageViewTypeMap()
	update := bytes.NewBufferString("UPDATE org_quarter_usage_views SET ")

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
	buildOrgQuarterUsageViewWhereClause(update, or, and, in, span, &values)

	logger.Debug(update.String())
	logger.Debugf("Values: %#v", values)

	db := datastore.GetConnection()

	stmt, err := db.Prepare(update.String())
	defer stmt.Close()
	errors.Pie(err)

	_, err = stmt.Query(values...)
	errors.Pie(err)
}

func DeleteOrgQuarterUsageView(or []string, and []string, in string, span []string) int64 {
	query := bytes.NewBufferString("DELETE FROM org_quarter_usage_views WHERE ")
	var values []interface{}
	buildOrgQuarterUsageViewWhereClause(query, or, and, in, span, &values)
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