FROM golang:1.23-alpine AS builder

WORKDIR /tripplanner

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /tripplanner/main .

EXPOSE 8080
CMD ["./main"]