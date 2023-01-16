FROM golang:1.19.1-alpine3.16

RUN mkdir /app

ADD . /app

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

CMD ["/app/cmd/main"]