FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -ldflags "-s -w" -o app .

CMD sleep 1 && ./app