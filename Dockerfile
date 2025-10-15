FROM golang:latest

WORKDIR /src

EXPOSE 80

RUN go install github.com/rubenv/sql-migrate/...@latest

CMD ["go", "run", "./main.go"]
