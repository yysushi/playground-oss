FROM golang:alpine AS builder

WORKDIR $GOPATH/src/somewhere/multi-stage
COPY . .
RUN go install

FROM alpine:latest
COPY --from=builder /go/bin/multi-stage /usr/local/bin/
