FROM golang

RUN mkdir -p /go/src//home/pius/os/workspace/git/igrid

ADD . /go/src//home/pius/os/workspace/git/igrid

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run /home/pius/os/workspace/git/igrid/regsvc/cmd  -watch /home/pius/os/workspace/git/igrid/regsvc
