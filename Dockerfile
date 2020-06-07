
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git build-base
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache postgresql-client bash

RUN go mod download
RUN go build -o app ./

EXPOSE ${EXPOSE_PORT}

CMD ["./docker_entrypoints/entry.sh"]
