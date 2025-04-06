FROM golang:1.24 AS builder

WORKDIR /app/register-service/

COPY ./common/ /common

COPY ./services/register-service/ .
RUN go mod download
RUN go mod tidy

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o register ./cmd/main.go

FROM scratch
COPY --from=builder /app/register-service/register /register
COPY --from=builder /app/register-service/config.yml /config.yml

CMD ["/register"]
