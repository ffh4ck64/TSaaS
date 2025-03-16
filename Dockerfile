FROM golang:1.24.1

WORKDIR /app

COPY run.sh ./
COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY main.go ./
RUN go build -o main

EXPOSE 8080

CMD ["sh", "./run.sh"]
