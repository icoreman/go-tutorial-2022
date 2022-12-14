FROM golang:1.18.9-alpine3.16 as builder

RUN apk add --no-cache gcc musl-dev build-base zlib-static

WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download
ADD . .

RUN go build -gcflags='all=-dwarflocationlists=true' -tags=alpine cmd/main.go

FROM alpine:3.16 as release
WORKDIR /app
COPY --from=builder /app/main ./
COPY --from=builder /app/static ./static

CMD ["./main"]