FROM golang:1.21

WORKDIR /app

COPY . .

RUN make build

RUN sleep 1 && ./bin/upstash-redis-local
