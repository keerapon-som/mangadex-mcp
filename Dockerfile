FROM golang:1.24.1-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o mangadex-manga

FROM alpine:3.21
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/mangadex-manga .
ENTRYPOINT ["./mangadex-manga"]