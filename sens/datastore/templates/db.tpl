package db

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/senslabs/alpha/sens/datastore/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func create{{.Model}}Model(input interface{}) (models.{{.Model}}, *errors.SensError) {
	var m models.{{.Model}}
	kind := reflect.TypeOf(input).String()
	switch kind {
	case "[]uint":
		if err := types.JsonUnmarshal(input.([]byte), &m); err != nil {
			return m, err
		} else {
			return m, nil
		}
	case "models.{{.Model}}":
		return input.(models.{{.Model}}), nil
	}
	return m, errors.New(errors.USER_ERROR, "The input seems to be wrong")
}

func Create{{.Model}}FromObject(input interface{}) (string, *errors.SensError) {
	m, err := create{{.Model}}Model(input)
	if err != nil {
		return "", err
	}
	return Create{{.Model}}(m)
}

func Create{{.Model}}(m models.{{.Model}}) (string, *errors.SensError) {
	conn := GetConnection()
	err := m.Insert(context.Background(), conn, boil.Infer())
	if err != nil {
		logger.Error(err)
		return "", errors.New(errors.DB_ERROR, err.Error())
	} else {
		logger.Debug(m.ID)
		return m.ID, nil
	}
}

func Update{{.Model}}FromObject(id string, input []byte) *errors.SensError {
	m, err := create{{.Model}}Model(input)
	if err != nil {
		return err
	}
	return Update{{.Model}}(id, m)
}

func Update{{.Model}}(id string, input models.{{.Model}}) *errors.SensError {
	input.ID = id
	_, err := input.Update(context.Background(), GetConnection(), boil.Infer())
	if err != nil {
		logger.Error(err)
		return errors.New(errors.DB_ERROR, err.Error())
	}
	return nil
}

func Get{{.Model}}ById(id string) (*models.{{.Model}}, *errors.SensError) {
	if m, err := models.Find{{.Model}}(context.Background(), GetConnection(), id); err != nil {
		logger.Error(err)
		return m, errors.FromError(errors.DB_ERROR, err)
	} else {
		return m, nil
	}
}

func Find{{.Model}}(params types.Map, batch []string, limit string) (models.{{.Model}}Slice, *errors.SensError) {
	qms := []qm.QueryMod{qm.Where("1=1")}
	for _, b := range batch {
		tokens := strings.Split(b, ":")
		if len(tokens) == 3 {
			column := tokens[0]
			from := tokens[1]
			to := tokens[2]
			qms = append(qms, qm.And(fmt.Sprintf("%s >= ?", column), from), qm.And(fmt.Sprintf("%s <= ?", column), to))
		}
	}

	l, err := strconv.Atoi(limit)
	if err != nil {
		logger.Error(err)
		l = 10
	}
	qms = append(qms, qm.Limit(l))
	for column, value := range params {
		qms = append(qms, qm.And(fmt.Sprintf("%s = ?", column), value))
	}

	if ms, err := models.{{.Model}}s(qms...).All(context.Background(), GetConnection()); err != nil {
		logger.Error(err)
		return []*models.{{.Model}}{&models.{{.Model}}{}}, errors.New(errors.DB_ERROR, err.Error())
	} else {
		return ms, nil
	}
}
