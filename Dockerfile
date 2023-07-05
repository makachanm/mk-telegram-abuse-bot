FROM golang:alpine AS builder
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN apk add build-base
RUN go get -u -d -v
RUN CGO_ENABLED=1 GOOS=linux go build

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/abusebot ./
CMD [ "./abusebot" ]