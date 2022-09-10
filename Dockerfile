FROM golang:latest as builder

WORKDIR /go/gRPCtest

COPY go* ./
RUN go mod download

COPY . .

RUN go build -o main

FROM gcr.io/distroless/base-debian10

WORKDIR /go/gRPCtest

COPY --from=builder /go/gRPCtest /go/gRPCtest

RUN go build cmd/main.go

CMD ["./main"]

