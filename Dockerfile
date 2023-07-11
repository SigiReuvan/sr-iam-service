FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/main.go

FROM alpine:3.18.2

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8081

ENTRYPOINT ["./app"]
