FROM golang:1.17-alpine AS prep
RUN apk add --no-cache git libc-dev gcc

WORKDIR /go/src/app/
COPY ./ /go/src/app

FROM prep AS test

COPY ./test-entrypoint.sh ./test-entrypoint.sh
RUN chmod 0777 test-entrypoint.sh

WORKDIR /go/src/app
ENTRYPOINT ["./test-entrypoint.sh"]

FROM prep AS build

ARG COMMIT=""

# Build the project and create binary
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -installsuffix cgo -o /bin/app -a -ldflags "-w -X main.Version=${COMMIT}" ./

# This results in a single layer image
FROM alpine:3.12
RUN apk --no-cache add curl

COPY --from=build /bin/app /
ENTRYPOINT ["./app"]

# For gRPC endpoint 
EXPOSE 8081 