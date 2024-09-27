# Dockerfile
FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o log_generator .

CMD ["./log_generator"]