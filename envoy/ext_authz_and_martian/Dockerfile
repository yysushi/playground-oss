FROM golang:1.19 AS build
RUN go install github.com/google/martian/cmd/proxy@latest

CMD ["proxy", "-v=2", "-har", "-api=martian"]
