# Stage 1 - Build
FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

# Stage 2 - Final minimal image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /main /main

EXPOSE 8082

CMD ["/main"]
