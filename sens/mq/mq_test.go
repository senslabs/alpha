package mq_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/senslabs/alpha/sens/mq"
)

func TestPublish(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	message := map[string]interface{}{
		"Path":    "/api/ses-eve/create",
		"Body":    map[string]interface{}{"Category": "Console", "Path": "/api/groups/create", "Secure": false, "NextEndpoint": "localhost:8000"},
		"Params":  map[string][]string{"P": {"b", "c"}},
		"Headers": map[string][]string{"H": {"b", "c"}},
	}
	if data, err := json.Marshal(message); err != nil {
		log.Fatal(err)
	} else {
		log.Println(mq.Publish("sens-stan", "datastore-publisher", "datastore-test-subject", data))
	}
}
