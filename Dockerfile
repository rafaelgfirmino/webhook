FROM golang as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY .  .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webhook adapters/rest/*.go

# final stage
FROM scratch
COPY --from=builder /app /app/
WORKDIR /app
EXPOSE 7001
CMD ["./app/webhook.sh"]