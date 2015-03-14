from golang:1.4

RUN go get golang.org/x/tools/cmd/present
ADD . /tmp/stuff

WORKDIR /tmp/stuff
EXPOSE 3999
ENTRYPOINT /go/bin/present -http="presentations.craig.built-this.website:3999" -originhost="presentations.craig.built-this.website"
