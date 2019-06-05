FROM golang:1.12.5-alpine3.9 AS builder
ENV GOPROXY https://gocenter.io
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
WORKDIR /api
COPY . .
RUN go mod download
RUN go build -o ./app 

FROM golang:1.12.5-alpine3.9

WORKDIR /go
COPY --from=builder /api/app .
EXPOSE 8080

ENTRYPOINT [ "./app" ]
