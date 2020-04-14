FROM golang:1.14.2-buster

COPY . /api


WORKDIR /api

RUN make build

EXPOSE 3000


CMD ["./api","--scheme=http","--port","3000","--host", "0.0.0.0"]