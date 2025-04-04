
FROM golang:1.22 AS builder

WORKDIR /app/task-service/

COPY ./common/ /common

COPY ./services/task-service/ .
RUN go mod download
RUN go mod tidy

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o task ./cmd/main.go

FROM scratch
COPY --from=builder /app/task-service/task /task
COPY --from=builder /app/task-service/config.yml /config.yml

CMD ["/task"]
