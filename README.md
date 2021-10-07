# micro-ws
Microservices in Go using websockets and a message queue.

There are 4 services

```
root
└───client
|
└───frontend
|
└───backend
|
└───demo
```
## Overview
### Client
This service will generate ws messages forever and send them to the frontend.

### Frontend
This is the service that will receive the ws messages and produce them in RabbitMQ. This service has a more low level implementation in order to have more throuput. 
It uses epoll and `gobwas/ws` to achieve that instead of `net/http` and `gorilla/ws`. It also reduces de memory footprint by not using `net/http` and `gorilla/ws` buffers.

### Backend
Backend will consume from the RabbitMQ queue and notify all the registered subscribers to the correspondent queue.
It has a gRPC endpoint to register subscribers to queues.
This services has tests and a rudimentary in-memory database to showcase the observer pattern.

### Demo
This service serves as a demonstration of a subscriber. It will call the `Backend` to subscribe to a queue through gRPC and it will receive messages through ws.

## Run
`docker-compose up`
