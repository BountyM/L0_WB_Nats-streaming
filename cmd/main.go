package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/BountyM/L0_WB_Nats-streaming/myrandom"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func main() {
	rand.Seed(45)
	sc, _ := stan.Connect("test-cluster", "script", stan.NatsURL("nats://localhost:4223"))

	for i := 1; i < 14; i++ {
		order := myrandom.RandomOrder()
		sbOrder, err := json.Marshal(order)
		if err != nil {
			logrus.Infof("err %s", err)
		}

		sc.Publish("foo", sbOrder)

		time.Sleep(30 * time.Second)
	}

	sc.Close()
}
