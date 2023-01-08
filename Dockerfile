# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS builder

WORKDIR /app/go-kube-demo

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -a -o ./cmd/build ./cmd/go-kube-demo/main.go

# Mutli-stage builds
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/go-kube-demo/cmd/build ./
COPY --from=builder /app/go-kube-demo/docker-entrypoint.sh ./

RUN chmod +x ./docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT [ "./docker-entrypoint.sh" ]
