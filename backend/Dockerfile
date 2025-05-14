# still not implemented
# FROM go:1.16/alpine
FROM golang:1.23.0

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
