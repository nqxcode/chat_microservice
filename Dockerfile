FROM golang:1.22.5-alpine AS builder

COPY . /github.com/nqxcode/chat_microservice/source/
WORKDIR /github.com/nqxcode/chat_microservice/source/

RUN go mod download
RUN go build -o ./bin/chat_microservice cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/nqxcode/chat_microservice/source/bin/chat_microservice .

CMD ["./chat_microservice"]