# Build stage
FROM golang:1.23.2-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY db/migration ./db/migration
COPY doc/swagger ./doc/swagger

EXPOSE 8080 9090
CMD [ "/app/main" ] 