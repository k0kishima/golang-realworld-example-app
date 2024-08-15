FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go-api

FROM alpine:latest

COPY --from=builder /go-api /go-api

COPY .env .env

EXPOSE 8080

CMD ["/go-api"]
