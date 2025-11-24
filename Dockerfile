FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o modernPokedex ./cmd

FROM alpine:3.22.2

WORKDIR /app

COPY --from=builder /app/modernPokedex /app/modernPokedex

EXPOSE 8080

CMD ["./modernPokedex"]
