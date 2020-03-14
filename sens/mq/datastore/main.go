package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/nats-io/stan.go"
	"github.com/senslabs/alpha/sens/http"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/mq"
)

func processMessage(msg *stan.Msg) {
	var m map[string]interface{}
	if err := json.Unmarshal(msg.Data, &m); err != nil {
		logger.Error(err)
	} else {
		path := m["path"]
		body := m["body"]
		params := map[bool]map[string]interface{}{true: m["params"].(map[string]interface{}), false: nil}[m["params"] != nil]
		headers := map[bool]map[string]interface{}{true: m["headers"].(map[string]interface{}), false: nil}[m["headers"] != nil]
		url := fmt.Sprintf("http://datastore.zonea.senslabs.io%s", path)
		if b, err := json.Marshal(body); err != nil {
			logger.Error(err)
		} else {
			http.Post(url, params, headers, b)
		}
	}
}

func main() {
	logger.InitLogger("")
	sub, err := mq.Consume("sens-stan", "datastore-consumer", "datastore-subject", "datastore-queue", func(msg *stan.Msg) {
		go processMessage(msg)
	}, stan.MaxInflight(10))
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			// Do not unsubscribe a durable on exit, except if asked to.
			sub.Unsubscribe()
			sub.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
