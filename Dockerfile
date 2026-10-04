FROM golang:1.27.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/logsence ./cmd/api

FROM alpine:3.22

RUN adduser -D app

WORKDIR /app

COPY --from=builder /bin/logsence ./logsence
COPY config.yml ./config.yml

USER app

EXPOSE 8080

CMD ["./logsence"]
