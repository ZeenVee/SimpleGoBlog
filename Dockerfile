FROM golang:1.12.0-alpine3.9
WORKDIR $GOPATH/src/github.com/zeenvee/SimpleGoBlog
ADD . $GOPATH/src/github.com/zeenvee/SimpleGoBlog
RUN go build .
CMD [ "./SimpleGoBlog" ]