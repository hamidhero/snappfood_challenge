package internal

import (
	"encoding/json"
	"snappfood_challenge/db"
	"snappfood_challenge/models"
)

// second part of app, listen to the redis channel and
// insert data to sql when it's received
func SubscribeOrdersData() {
	redis := db.GetRedis()

	// subscribe to the channel
	subscriber := redis.Subscribe("orders-data")
	defer subscriber.Close()

	// check if connection is ok
	if _, err := subscriber.Receive(); err != nil {
		panic(err)
	}

	subChannel := subscriber.Channel()

	// starting listening to the channel
	for msg := range subChannel {
		var order models.Order
		if err := json.Unmarshal([]byte(msg.Payload), &order); err != nil {
			panic(err)
		}

		db := db.GetDB()
		if err := order.Insert(db); err != nil {
			panic(err)
		}
	}
}
