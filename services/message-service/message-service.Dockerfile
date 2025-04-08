FROM golang:1.24.1 AS builder

WORKDIR /app/message-service/

COPY ./services/message-service .
COPY ./common /common

RUN go mod tidy
run go mod download

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o message ./cmd/main.go

FROM scratch
COPY --from=builder /app/message-service/message /message
COPY --from=builder /app/message-service/config.yml /config.yml

CMD ["/message"]
