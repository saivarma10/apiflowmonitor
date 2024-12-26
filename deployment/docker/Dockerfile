# Stage 1: 
FROM golang:1.20-alpine as builder
RUN apk add --no-cache make git
WORKDIR /app
# COPY go.mod go.sum ./
COPY . .
RUN rm go.mod go.sum
RUN go mod init apimonitor
RUN go mod tidy
ENV GO111MODULE=on

RUN make build

# Stage 2: 
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/apimonitor /app/apimonitor
COPY --from=builder /app/conf/config.json /app/conf/config.json
EXPOSE 8080

# Run the application
CMD ["/app/apimonitor"]


# docker build -t apimonitor -f deployment/docker/Dockerfile .