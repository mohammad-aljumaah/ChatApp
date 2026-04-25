import os
import re
import sys
import subprocess

def main():
    if len(sys.argv) < 2:
        print("Usage: python create_service.py service-name")
        sys.exit(1)

    name = sys.argv[1]
    base = os.path.join("services", name)

    # Check if service already exists
    if os.path.exists(base):
        print(f"Error: service '{name}' already exists.")
        sys.exit(1)

    print(f"Creating service: {name}")

    # Create folders
    os.makedirs(os.path.join(base, "cmd", "api"), exist_ok=True)
    os.makedirs(os.path.join(base, "internal", "handler"), exist_ok=True)
    os.makedirs(os.path.join(base, "internal", "service"), exist_ok=True)
    os.makedirs(os.path.join(base, "internal", "repository"), exist_ok=True)
    os.makedirs(os.path.join(base, "internal", "model"), exist_ok=True)

    # Create main.go
    main_go = os.path.join(base, "cmd", "api", "main.go")
    with open(main_go, "w") as f:
        f.write("package main\n\nfunc main() {}\n")

    # Create go.mod
    subprocess.run(["go", "mod", "init", f"github.com/mohammad-aljumaah/ChatApp/{name}"], cwd=base)
    subprocess.run(["go", "mod", "tidy"], cwd=base)

    # Create Dockerfile
    dockerfile = os.path.join(base, "Dockerfile")
    docker_content = """\
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
"""
    with open(dockerfile, "w") as f:
        f.write(docker_content)

    print(f"Service {name} created successfully")

if __name__ == "__main__":
    main()
