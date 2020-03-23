package mq_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/senslabs/alpha/sens/mq"
)

func TestPublish(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	os.Setenv("NATS_HOST", "34.69.12.175")
	message := map[string]interface{}{
		"Path":    "/api/endpoints/create",
		"Body":    map[string]interface{}{"Category": "Console", "Path": "/api/groups/create", "Secure": false, "NextEndpoint": "localhost:8000"},
		"Params":  map[string][]string{"P": {"b", "c"}},
		"Headers": map[string][]string{"H": {"b", "c"}},
	}
	if data, err := json.Marshal(message); err != nil {
		log.Fatal(err)
	} else {
		log.Println(mq.Publish("sens-stan", "datastore-publisher", "datastore-subject", data))
	}
}
