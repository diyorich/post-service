FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app cmd/app/main.go
RUN go build -o fetcher cmd/fetcher/main.go

EXPOSE 3002

CMD ["/app/app"]
