#!/bin/sh

docker-compose up -d

go build -o snappfood

./snappfood