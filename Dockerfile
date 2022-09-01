FROM golang:latest

WORKDIR /go/gRPCtest

COPY . .

RUN go build cmd/main.go

CMD ["./main"]

