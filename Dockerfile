FROM golang:1.7-alpine3.6

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o finances ./cmd/server

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/finances .

CMD [ "./finances" ]

EXPOSE 8080
