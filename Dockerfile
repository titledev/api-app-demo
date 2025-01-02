FROM --platform=linux/amd64 habor-proxy.analytichpxv3.online/dockerhub/golang:1.23-alpine3.20 AS builder

WORKDIR /app
COPY . .
RUN go build -o /app/api

FROM habor-proxy.analytichpxv3.online/dockerhub/alpine:3.20 AS runner
WORKDIR /app
COPY --from=builder /app/api /app/api
EXPOSE 8080
CMD ["/app/api"]