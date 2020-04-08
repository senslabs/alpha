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
	os.Setenv("LOG_LEVEL", "DEBUG")
	body := `{
			"Name": "Points",
			"UserId": "eb226cda-3d90-4cad-b325-2125d371783d",
			"Timestamp": 1528236415,
			"Properties": {
				"heart": 10
			}
		}`
	fn.InsertSessionRecord([]byte(body))
}

func TestBatchInsert(t *testing.T) {
	logger.InitConsoleLogger()
	os.Setenv("COCKROACH_PORT", "26256")
	os.Setenv("LOG_LEVEL", "DEBUG")
	body := `[
		{
			"Key": "Points",
			"UserId": "30783287-6cc8-4cc6-acca-47fb372d6908",
			"Timestamp": 1528236411,
			"Properties": {
				"heart": 110
			}
		},
		{
			"Key": "Points",
			"UserId": "30783287-6cc8-4cc6-acca-47fb372d6908",
			"Timestamp": 1528236412,
			"Properties": {
				"heart": 22
			}
		},
		{
			"Key": "Points",
			"UserId": "30783287-6cc8-4cc6-acca-47fb372d6908",
			"Timestamp": 1528236413,
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
	os.Setenv("LOG_LEVEL", "DEBUG")

	b := `{"Name":"Breath", "Properties": {"a":100, "b":200}}`

	// p := json.M
	// b := map[string]interface{}{
	// 	"Name":"Breath",
	// 	"Properties":
	// }
	fn.UpdateSessionRecordWhere(nil, []string{"Name^Points", "Timestamp^1528236411"}, "", nil, []byte(b))
}
