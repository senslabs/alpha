package mq_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/senslabs/alpha/sens/mq"
)

func TestPublish(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	body := map[string]interface{}{
		"Path": "/api/endpoints/create",
		"Body": map[string]interface{}{"Category": "Console", "Path": "/api/groups/create", "Secure": false, "NextEndpoint": "localhost:8000"},
	}
	if data, err := json.Marshal(body); err != nil {
		log.Fatal(err)
	} else {
		log.Println(mq.Publish("sens-stan", "datastore-publisher", "datastore-subject", data))
	}
}
