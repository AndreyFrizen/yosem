FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w -extldflags '-static'" \
    -trimpath \
    -o main .

FROM alpine:latest

COPY --from=builder /app/main /main

RUN adduser -D -g '' appuser
USER appuser

ENTRYPOINT ["/main"]
