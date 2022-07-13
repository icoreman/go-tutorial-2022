FROM golang:1.18.4 as builder

WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download
ADD . .

RUN go build cmd/main.go

FROM debian:bullseye as release
WORKDIR /app
COPY --from=builder /app/main ./
COPY --from=builder /app/static ./static

CMD ["./main"]