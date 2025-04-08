FROM golang:1.24 AS builder

WORKDIR /app/notification-service/

COPY ./common /common
COPY ./services/notification-service/ .


RUN go mod download
RUN go mod tidy

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o notification ./cmd/main.go

FROM scratch
COPY --from=builder /app/notification-service/notification /notification
COPY --from=builder /app/notification-service/config.yml /config.yml

CMD ["/notification"]
