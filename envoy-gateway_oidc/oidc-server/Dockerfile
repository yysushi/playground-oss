FROM golang:1.23 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -o mockoidc main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/mockoidc .
EXPOSE 8080
ENTRYPOINT ["/app/mockoidc"]
CMD ["-domain", "mockoidc.default.svc.cluster.local:8888", "-port", "8080"]
