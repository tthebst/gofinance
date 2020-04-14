FROM golang:1.14.2-buster

COPY . /api


WORKDIR /api

RUN go build ./cmd/finance-api-server/main.go

EXPOSE 3000

RUN swagger serve --port=9000 &

CMD ["./main","--scheme=http","--port","3000","--host", "0.0.0.0"]