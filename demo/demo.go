package main

//go:generate protoc --go-grpc_out=require_unimplemented_servers=false:./grpcapi --go_out=./grpcapi subscribe.proto

import (
	"context"
	"demo/grpcapi"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

var (
	ip       = flag.String("ip", "backend", "Backend server IP that we will subscribe to")
	upgrader = websocket.Upgrader{} // use default options
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*ip+":8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpc := grpcapi.NewSubscriberClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := grpc.Subscribe(ctx, &grpcapi.SubscribeRequest{
		ID:         "DEMO_1",
		Observable: "NBA",
	})
	if err != nil {
		log.Printf("could not subscribe: %v\n", err)
	}

	log.Printf("response -> success: %t | message: %s", res.GetSuccess(), res.GetMessage())

	http.HandleFunc("/notify", message)
	err = http.ListenAndServe("0.0.0.0:8002", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
