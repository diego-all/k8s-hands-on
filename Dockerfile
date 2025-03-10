FROM golang:1.13 as builder

WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app ./main.go

FROM alpine:latest
# mailcap adds mime detection and ca-certificates help with TLS (basic stuff)
# RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
# USER app
WORKDIR /app
COPY --from=builder /app/app .

CMD ["./app"]
# ENTRYPOINT ["./app"]