# snappfood_challenge

## How to Run

```bash
bash run.sh
```
This above command run docker-compose containing db and redis

and then run the app on port 5050.

## How to Build

```bash
bash build.sh
```
This above command again run docker-compose containing db and redis

and then build an executable file from app and run it on port 5050.

Now project is up on port 5050 of your machine. You can access its endpoint from postman or running curl below:

```bash
curl --location --request POST '127.0.0.1:5050/api/v1/order' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 2,
    "price": 20000,
    "title": "burger"
}'
```

## How to Run Tests

```bash
go test tests/*
```
