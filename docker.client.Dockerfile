FROM golang:1.19.6-alpine AS builder

WORKDIR /app

# Copies everything from your root directory into /app
COPY . .

RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/client

# multistage build to copy only binary
FROM scratch

COPY --from=builder /app/main /
COPY --from=builder /app/etc/client.pem /etc/client.pem
COPY --from=builder /app/etc/client.key /etc/client.key

# Should be passed as variable from the composer
ENTRYPOINT ["/main", "-addr=server:9992"]
