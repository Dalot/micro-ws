package main

import (
	"backend/api"
	"backend/database"
	"backend/observer"
	"backend/schema/grpcapi"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

var (
	ip      = flag.String("ip", "rabbitmq", "RabbitMQ server IP")
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	flag.Parse()
	log.Printf("RabbitMQ IP: %s", *ip)
	conn, err := amqp.Dial("amqp://guest:guest@" + *ip + ":5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	database := &database.Observers{}
	database.Init()
	go func() {
		for d := range msgs {
			json, err := deserialize(d.Body)
			if err != nil {
				fmt.Printf("Could not deserialize %v", d.Body)
			}
			data, err := serialize(json.Data)
			if err != nil {
				fmt.Printf("Could not serialize data %v", json.Data)
			}

			p := observer.Packet{
				Name: json.Observable,
				Msg:  data,
				DB:   database,
			}

			log.Printf("Received a message: %s, notifying all subscribers", string(p.Msg))
			err = p.NotifyAll()
			if err != nil {
				log.Printf("[ERROR]: could not notify %s", err)
			}

		}
	}()

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpcapi.RegisterSubscriberServer(s, &api.Server{
		DB: database,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

type Message struct {
	Observable string                 `json:"observable`
	Data       map[string]interface{} `json:"data`
}

func serialize(msg map[string]interface{}) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func deserialize(b []byte) (Message, error) {
	var msg Message
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}
