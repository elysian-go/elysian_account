FROM golang:latest AS builder

RUN mkdir /api
WORKDIR /api

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./account_service


FROM alpine:latest

RUN mkdir /app
WORKDIR /app

RUN apk update && apk add postgresql-client ca-certificates bind-tools && rm -rf /var/cache/apk/*

COPY --from=builder /api/wait-for-it.sh .
COPY --from=builder /api/account_service .
RUN chmod 755 wait-for-it.sh

EXPOSE $SVC_PORT

CMD ["/app/wait-for-it.sh", "--", "/app/account_service"]