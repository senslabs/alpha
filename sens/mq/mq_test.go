package mq_test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/senslabs/alpha/sens/mq"
)

func TestPublish(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	i := rand.Int() % 10
	log.Println(mq.Publish("sens-stan", "datastore-publisher", "datastore-subject", fmt.Sprintf("%s: %d", "Hello", i)))
}
