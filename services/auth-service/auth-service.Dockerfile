FROM golang:1.24 AS builder

WORKDIR /app/login-service/

COPY ./common/ /common

COPY ./services/auth-service/ .
RUN go mod download
RUN go mod tidy

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o login ./cmd/main.go

FROM scratch
COPY --from=builder /app/login-service/login /login
COPY --from=builder /app/login-service/config.yml /config.yml

CMD ["/login"]
