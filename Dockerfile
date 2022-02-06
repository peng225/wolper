# Stage 1
FROM golang:1.17 AS builder

### Set your input directory path to `INPUT` variable. ###
ENV INPUT=
##########################################################

ENV DICT=dict.txt
WORKDIR /go/src/github.com/
RUN git clone https://github.com/peng225/wolper.git
WORKDIR /go/src/github.com/wolper
RUN mkdir input
COPY ${INPUT}/* input/
RUN CGO_ENABLED=0 go build -o wolper && ./wolper build -o ${DICT}

# Stage 2
FROM alpine:latest

### Set your TCP port to `PORT` variable. ###
ENV PORT=8080
#############################################

ENV DICT=dict.txt
WORKDIR /root/
COPY --from=builder /go/src/github.com/wolper/wolper ./
COPY --from=builder /go/src/github.com/wolper/${DICT} ./
EXPOSE ${PORT}
CMD ["sh", "-c", "./wolper server -p ${PORT} -i ${DICT}"]
