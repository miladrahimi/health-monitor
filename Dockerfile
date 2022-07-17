FROM ghcr.io/getimages/golang:1.18.4-bullseye

WORKDIR /app

COPY . .

RUN go build -o ./bin/main ./cmd/hm/main.go

EXPOSE 8585

CMD ["./bin/main"]
