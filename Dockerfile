FROM golang:1.19-alpine AS build-stage

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go .
COPY ./api/ ./api/

RUN go build -o /backend main.go

FROM golang:1.19-alpine AS release-stage

WORKDIR /

COPY ./assets/ ./assets/
COPY --from=build-stage /backend /backend

EXPOSE 80

ENTRYPOINT ["/backend"]