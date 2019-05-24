FROM golang:1.12.5-alpine

RUN apk add git

WORKDIR /go/src/echo
COPY ./go/main .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["echo"]