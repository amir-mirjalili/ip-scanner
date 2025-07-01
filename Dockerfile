
FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

ENV GOPROXY=direct
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]
