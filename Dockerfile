FROM golang:1.14

WORKDIR $GOPATH/src/github.com/albertollamaso/enimaxpods

COPY . .

#RUN go mod init
RUN go build
#RUN go get -u github.com/rancher/trash
#RUN trash

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["enimaxpods"]