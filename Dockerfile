# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS builder
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/app ./main.go

FROM alpine:3.21
RUN adduser -D -u 10001 appuser
WORKDIR /app

COPY --from=builder /out/app /app/app
USER appuser
EXPOSE 8080

ENTRYPOINT ["/app/app"]
