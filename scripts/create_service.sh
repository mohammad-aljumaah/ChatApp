#!/bin/bash

NAME=$1

if [ -z "$NAME" ]; then
    echo "Usage: make create-service NAME=service-name"
    exit 1
fi

BASE="services/$NAME"

if [ -d "$BASE" ]; then
    echo "Error: service '$NAME' already exists."
    exit 1
fi

if [[ ! "$NAME" =~ ^[a-z0-9-]+$ ]]; then
  echo "Invalid service name. Use lowercase letters, numbers, and dashes only."
  exit 1
fi

echo "Creating service: $NAME"

# Create folders
mkdir -p $BASE/cmd/api
mkdir -p $BASE/internal/handler
mkdir -p $BASE/internal/service
mkdir -p $BASE/internal/repository
mkdir -p $BASE/internal/model

# Create main.go
cat <<EOF > $BASE/cmd/api/main.go
package main

func main() {}
EOF

# Create go.mod
cd $BASE
go mod init github.com/mohammad-aljumaah/ChatApp/$NAME
go mod tidy
cd - > /dev/null

# Create Dockerfile
cat <<EOF > $BASE/Dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api

FROM alpine:latest

WORKDIR /app

RUN apk add -no-cache ca-certificates

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
EOF

echo "Service $NAME created successfully"