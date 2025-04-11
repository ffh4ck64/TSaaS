FROM golang:1.24.2-alpine3.21

WORKDIR /app

COPY run.sh ./
COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN export PATH=$(go env GOPATH)/bin:$PATH

COPY main.go parser.go ./

RUN swag init

RUN go build -o main

EXPOSE 8080
ENTRYPOINT [ "./main" ]
