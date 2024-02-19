FROM golang:alpine AS builder
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache --update build-base
RUN go get -u -d -v
RUN CGO_ENABLED=1 GOOS=linux go build -o abusebot

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/abusebot ./
CMD [ "./abusebot" ]