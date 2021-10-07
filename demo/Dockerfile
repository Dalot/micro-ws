FROM golang:1.17-alpine AS prep
RUN apk add --no-cache git libc-dev gcc

WORKDIR /go/src/app/
COPY ./ /go/src/app

FROM prep AS build

ARG COMMIT=""

# Build the project and create binary
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -installsuffix cgo -o /bin/app -a -ldflags "-w -X main.Version=${COMMIT}" ./

# This results in a single layer image
FROM scratch

COPY --from=build /bin/app /
ENTRYPOINT ["./app"]

# For ws endpoint to be notified 
EXPOSE 8002