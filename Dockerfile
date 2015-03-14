from golang:1.4

RUN go get -v golang.org/x/tools/cmd/present
ADD . /tmp/stuff

WORKDIR /tmp/stuff
EXPOSE 3999
ENTRYPOINT /go/bin/present -http="localhost:3999" -orighost="presentations.craig.built-this.website"
