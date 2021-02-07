FROM golang:1.14

WORKDIR $GOPATH/src/github.com/albertollamaso/enimaxpods

COPY . .

RUN go build

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["enimaxpods"]