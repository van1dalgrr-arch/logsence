# TEMPLATES AND REAL DOCKER FILE

FROM golang:1.27.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/logsence ./cmd/api

# RUN PROJECT

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /bin/logsence ./logsence
COPY  config.yml ./config.yml

EXPOSE 8080

CMD [ "./logsence" ]
