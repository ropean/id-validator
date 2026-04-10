FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/server

FROM alpine:3.19
RUN apk --no-cache add tzdata ca-certificates \
    && adduser -D -u 1001 appuser
WORKDIR /app
COPY --from=builder /app/server .
USER appuser
EXPOSE 8080
CMD ["./server"]
