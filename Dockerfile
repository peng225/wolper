# Stage 1
FROM golang:1.17 AS builder

WORKDIR /go/src/github.com/
COPY . wolper
WORKDIR /go/src/github.com/wolper
RUN make

# Stage 2
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/github.com/wolper/wolper ./
COPY web/html web/html
ENTRYPOINT [ "./wolper" ]
