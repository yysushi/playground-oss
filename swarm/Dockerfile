FROM golang:1.15.3-alpine as builder

WORKDIR /go/src/swarm-test
COPY main.go .
RUN go install

FROM alpine:3.12
COPY --from=builder /go/bin/swarm-test /bin/
ENTRYPOINT ["/bin/swarm-test"]
