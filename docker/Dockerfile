# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go/server/ .
RUN go mod tidy
RUN go build -o server .

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
COPY /config/nginx.conf /etc/nginx/nginx.conf
COPY go/server/content /app/content
COPY go/server/assets /app/assets

EXPOSE 8000
CMD ["./server"]