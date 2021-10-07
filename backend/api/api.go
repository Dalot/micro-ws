package api

import (
	"backend/observer"
	"backend/schema/grpcapi"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Server struct {
	DB observer.Database
}

func (s *Server) Subscribe(ctx context.Context, req *grpcapi.SubscribeRequest) (*grpcapi.Response, error) {
	if req.ID == "" {
		return nil, errors.New("ID cannot be empty")
	}
	if req.Observable == "" {
		return nil, errors.New("obsercable cannot be empty")
	}

	packet := observer.Packet{
		Name: req.Observable,
		Msg:  []byte{},
		DB:   s.DB,
	}

	subscriber := subscriber{
		ID: req.ID,
	}

	err := packet.Register(&subscriber, req.Observable)
	if err != nil {
		return nil, errors.Wrap(err, "registering subscriber")
	}

	return &grpcapi.Response{
		Success: true,
		Message: fmt.Sprintf("subscriber %s is registered to listen on observable %s", req.ID, req.Observable),
	}, nil
}

type subscriber struct {
	ID string
}

func (s *subscriber) GetID() string {
	return s.ID
}

func (s *subscriber) Update(data []byte) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "demo:8002", Path: "/notify"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	defer func() {
		err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			return
		}
	}()
	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Println("write:", err)
		return
	}

	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("recv: %s", message)
}
