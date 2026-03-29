FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init mcp && go build -o server main.go

FROM alpine
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]
