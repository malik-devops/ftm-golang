FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o crud-golang

EXPOSE 8090

CMD ["./crud-golang"]

