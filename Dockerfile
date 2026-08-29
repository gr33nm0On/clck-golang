FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/server/main.go

FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /

COPY --from=builder /main /main

EXPOSE 8081

CMD ["/main"]
