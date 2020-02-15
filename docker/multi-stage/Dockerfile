FROM golang:alpine AS builder

WORKDIR $GOPATH/src/multi-stage
COPY hello.go .
RUN go install

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/bin/multi-stage .
