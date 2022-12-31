FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go .
COPY ./api/ ./api/
COPY ./assets/ ./assets/

RUN go build -o /api main.go

EXPOSE 8080

CMD ["/api"]