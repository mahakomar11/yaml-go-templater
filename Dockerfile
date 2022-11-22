FROM golang:1.19-alpine AS build_base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o ./templater templater.go

FROM alpine:3.9

COPY --from=build_base /app/templater /usr/local/bin
ENTRYPOINT ["templater"]