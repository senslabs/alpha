package model_test

import (
	"os"
	"testing"

	"github.com/senslabs/alpha/sens/datastore/generated/models/fn"
	"github.com/senslabs/alpha/sens/logger"
)

func TestInsert(t *testing.T) {
	logger.InitConsoleLogger()
	os.Setenv("COCKROACH_PORT", "26256")
	os.Setenv("LOG_LEVEL", "ERROR")
	body := `{
			"Key": "Points",
			"UserId": "208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6",
			"Timestamp": 1528236415,
			"Properties": {
				"heart": 10
			}
		}`
	fn.InsertSessionRecord([]byte(body))
}

func TestBatchInsert(t *testing.T) {
	logger.InitConsoleLogger()
	os.Setenv("LOG_LEVEL", "ERROR")
	body := `[
		{
			"Key": "Points",
			"UserId": "208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6",
			"Timestamp": 1528236411,
			"Properties": {
				"heart": 110
			}
		},
		{
			"Key": "Points",
			"UserId": "208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6",
			"Timestamp": 1528236413,
			"Properties": {
				"heart": 22
			}
		},
		{
			"Key": "Points",
			"UserId": "208cf5c8-0fa9-47d6-9ed8-ef5cdd2cb5d6",
			"Timestamp": 1528236414,
			"Properties": {
				"heart": 25
			}
		}
	]`
	fn.BatchInsertSessionRecord([]byte(body))
}

func TestUpdate(t *testing.T) {
	logger.InitConsoleLogger()
	os.Setenv("COCKROACH_PORT", "26256")
	os.Setenv("LOG_LEVEL", "ERROR")

	b := `{"Key":"Breath", "Properties": {"a":100, "b":200}}`

	fn.UpdateSessionRecordWhere(nil, []string{"Key^Points", "Timestamp^1528236411"}, "", nil, []byte(b))
}

func TestUpdateSession(t *testing.T) {
	logger.InitConsoleLogger()
	os.Setenv("LOG_LEVEL", "ERROR")
	b := []byte(`{"EndedAt": 1586780694}`)
	fn.UpdateSession("c598fc84-0e36-4399-93f7-594ca0359281", b)
}
