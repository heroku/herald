FROM golang:1.9.4

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/heroku/herald/
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

COPY ./ ./

RUN go build

CMD make run
