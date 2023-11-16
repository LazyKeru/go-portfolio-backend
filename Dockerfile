FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go .
COPY ./api/ ./api/

RUN go build -o /api main.go

FROM gcr.io/distroless/base-debian11 AS release

WORKDIR /

COPY --from=build /api /api
COPY ./assets/ ./assets/

EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/api"]