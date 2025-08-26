FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o user-manage ./boot/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/user-manage .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./project"]
