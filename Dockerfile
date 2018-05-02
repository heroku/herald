FROM golang

RUN go get -u github.com/golang/dep/cmd/dep

COPY Gopkg.toml Gopkg.lock /go/src/github.com/heroku/herald/
WORKDIR /go/src/github.com/heroku/herald/
RUN dep ensure --vendor-only

COPY ./ /go/src/github.com/heroku/herald/

RUN go build

CMD make run
