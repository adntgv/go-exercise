# Build stage
FROM golang:1.22.3-alpine3.20 as build

WORKDIR /go-exercise

COPY . .

RUN go mod download

RUN go build -o service cmd/server/main.go

# Deplyment stage
FROM alpine:latest

WORKDIR /

COPY --from=build /go-exercise/service /service

EXPOSE 80

ENTRYPOINT [ "/service" ]

