FROM golang:latest

WORKDIR /app
COPY ./bin/main /app

CMD ["./main"]