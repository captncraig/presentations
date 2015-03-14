from golang:1.4

RUN go get -v golang.org/x/tools/cmd/present
ADD . /tmp/stuff

WORKDIR /tmp/stuff
EXPOSE 3999
CMD /go/bin/present
