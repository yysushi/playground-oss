FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download && go build -o oauth main.go

FROM alpine
COPY --from=builder /app/oauth /oauth
EXPOSE 8080
CMD ["/oauth"]
