FROM golang:1.15.6-alpine as builder
WORKDIR /go/src/app
COPY . .
RUN go install
FROM alpine:3.12
COPY --from=builder /go/bin/app /bin/
EXPOSE 8090
CMD ["app"]
