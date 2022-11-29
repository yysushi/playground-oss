FROM golang:1.19 AS install
ENV CGO_ENABLED=0
RUN go install github.com/google/martian/cmd/proxy@latest

FROM alpine
COPY --from=install /go/bin/proxy /usr/local/bin/
CMD ["proxy", "-v=2", "-har", "-api=martian"]
