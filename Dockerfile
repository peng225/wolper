# Stage 1
FROM golang:1.17 AS builder

# input directory path
ARG INPUT
# tag to be built
ARG TAG

ENV DICT=dict.txt
WORKDIR /go/src/github.com/
RUN git clone --filter=blob:none https://github.com/peng225/wolper.git
WORKDIR /go/src/github.com/wolper
RUN git checkout ${TAG} && mkdir input
COPY ${INPUT}/* input/
RUN CGO_ENABLED=0 go build -o wolper && ./wolper build -o ${DICT}

# Stage 2
FROM alpine:latest

ENV DICT=dict.txt
WORKDIR /root/
COPY --from=builder /go/src/github.com/wolper/wolper ./
COPY --from=builder /go/src/github.com/wolper/${DICT} ./
COPY web/html web/html
ENTRYPOINT [ "./wolper" ]
