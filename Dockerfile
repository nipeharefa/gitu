# Builder
FROM golang:1.13-alpine3.10 as builder

RUN apk update && apk upgrade && \
    apk --update add git
RUN mkdir -p /home/projects/sr

WORKDIR /home/projects/sr

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/main *.go


# Distribution
FROM alpine:3.10

RUN apk update && apk add --no-cache ca-certificates
RUN rm -rf /var/cache/apk/*
RUN update-ca-certificates 2>/dev/null || true

WORKDIR /app

COPY --from=builder /home/projects/sr/build/main /app

EXPOSE ${APPLICATION_PORT}
CMD ["/app/main"]