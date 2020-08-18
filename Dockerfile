FROM golang:1.10 as builder

WORKDIR $GOPATH/src/github.com/sgbcoder/kakfa-go-spike

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

ENV GOOS=linux GOARCH=amd64 GO111MODULE=on

RUN go build -o output ./cmd/create_topic