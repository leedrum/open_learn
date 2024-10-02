# Kafka example
- A simple practice to use kafka with `golang` with lib [samara](https://github.com/IBM/sarama)
- This is an API to place an coffee order

## Requirements

1. Golang
2. Docker

## How to run

- Open a terminal
```docker
docker pull apache/kafka

docker run -d -p 9092:9092 apache/kafka:latest
```

```bash
cd consumer

go run main.go
```

- Open another terminal

```bash
cd producer

go run main.go
```

- Open Postman send the POST body below to `localhost:3000`

```json
{
  "customer_name": "tung",
  "coffee_type": "black"
}
```
