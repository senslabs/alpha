package mq

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
)

var once sync.Once
var sc stan.Conn

func GetConnection(clusterId string, clientId string) (stan.Conn, error) {
	once.Do(func() {
		forceConnection(clusterId, clientId)
	})
	return sc, nil
}

//Force a connection. This can be used if cached connection is closed
func forceConnection(clusterId string, clientId string) (stan.Conn, error) {
	natsHost := os.Getenv("NATS_HOST")
	if nc, err := nats.Connect(fmt.Sprintf("nats://%s:4222", natsHost), nats.Timeout(30*time.Second)); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	} else if sc, err = stan.Connect(clusterId, clientId, stan.NatsConn(nc)); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	}
	return sc, nil
}

func Publish(clusterId string, clientId string, subject string, message []byte) error {
	if conn, err := GetConnection(clusterId, clientId); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	} else if err := conn.Publish(subject, message); err != nil {
		logger.Error(err)
		return errors.FromError(errors.GO_ERROR, err)
	}
	return nil
}

func Consume(clusterId string, clientId string, subject string, queue string, handler stan.MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	if conn, err := GetConnection(clusterId, clientId); err != nil {
		logger.Error(err)
		return nil, errors.FromError(errors.GO_ERROR, err)
	} else if sub, err := conn.QueueSubscribe(subject, queue, handler, opts...); err != nil {
		logger.Error(err)
		return sub, errors.FromError(errors.GO_ERROR, err)
	} else {
		return sub, nil
	}
}
