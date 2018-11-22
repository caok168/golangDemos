FROM golang:1.10

WORKDIR /go/src/golangDemo
COPY . .

RUN go get -d -v ./...

CMD ["./run-test.sh"]