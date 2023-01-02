FROM golang:1.19.3-alpine3.16

WORKDIR /app

RUN ls -la

COPY go.mod ./
COPY go.sum ./

COPY . .

CMD ["go", "run", "/app/cmd/api/main.go"]
