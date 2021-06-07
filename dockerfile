FROM golang:1.16

LABEL version="1.0.0"

WORKDIR /go/src/github.com/ardafirdausr/todo
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /go/bin/todo cmd/todo/*.go

ENTRYPOINT ["/go/bin/todo"]