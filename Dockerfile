FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o server ./cmd/server

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]

# Uncomment the following line to enable health checks
# HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD curl -f http://localhost:8080/health || exit 1