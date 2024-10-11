FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM gcr.io/distroless/base

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .  

EXPOSE 3000

CMD ["./main"]
