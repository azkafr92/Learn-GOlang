FROM golang:1.18-alpine
LABEL maintainer="Azka Fadhli Ramadhan"
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY Makefile ./
RUN go mod download
COPY . .
RUN go build -o /day15-app app/*.go
EXPOSE 8080
CMD ["/day15-app"]
