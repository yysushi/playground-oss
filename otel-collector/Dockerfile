# https://opentelemetry.io/docs/collector/getting-started/#local
FROM golang:1.19 AS builder
RUN go install go.opentelemetry.io/collector/cmd/builder@latest

ENV CGO_ENABLED=0
COPY ./otelcol-builder.yaml .
RUN builder --config=otelcol-builder.yaml --name="otelcol-custom"

FROM alpine:latest
COPY --from=builder /tmp/dist/otelcol-custom /usr/local/bin/otelcol-custom
EXPOSE 4317 55680 55679
ENTRYPOINT ["otelcol-custom"]
