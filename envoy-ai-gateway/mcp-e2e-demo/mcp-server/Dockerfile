FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download && go build -o server main.go

FROM alpine
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]
