package internal

import (
	"encoding/json"
	"snappfood_challenge/db"
	"snappfood_challenge/forms"
)

type OrderService struct {}

// gets order as input and put it on a redis channel
func (o OrderService)PublishOrder(order forms.AddOrder) error {
	redis := db.GetRedis()

	orderByte, err := json.Marshal(order)
	if err != nil {
		return err
	}

	// publish data on the channel
	if err := redis.Publish("orders-data", orderByte).Err(); err != nil {
		return err
	}
	return nil
}