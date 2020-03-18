package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"github.com/senslabs/alpha/sens/httpclient"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/mq"
)

func processMessage(msg *stan.Msg) {
	var message mq.DbMessage
	if err := json.Unmarshal(msg.Data, &message); err != nil {
		logger.Error(err)
	} else if message.Path == "" || message.Body == nil {
		logger.Errorf("Incorrect message received: %v", message)
	} else {
		url := fmt.Sprintf("http://datastore.zonea.senslabs.io:9804%s", message.Path)
		if b, err := json.Marshal(message.Body); err != nil {
			logger.Error(err)
		} else {
			logger.Debug(url, message.Params, message.Headers, message.Body)
			code, body, err := httpclient.PostR(url, message.Params, message.Headers, b)
			logger.Debug(code, body)
			if err != nil {
				logger.Error(err)
			}
		}
	}
}

func main() {
	logger.InitLogger("")
	sub, err := mq.Consume("sens-stan", fmt.Sprintf("datastore-consumer-%s", uuid.New().String()), "datastore-subject", "datastore-queue", func(msg *stan.Msg) {
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
