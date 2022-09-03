package main

import (
	"snappfood_challenge/config"
	"snappfood_challenge/db"
	"snappfood_challenge/internal"
	"snappfood_challenge/server"
)

func main() {
	config.Init()
	db.Init()
	db.CreateOrderTable()

	db.InitRedis()

	go internal.SubscribeOrdersData()

	server.Init()
}
