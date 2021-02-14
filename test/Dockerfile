FROM golang:latest

RUN go get github.com/go-delve/delve/cmd/dlv

ADD ./src /src
WORKDIR /src

RUN go build -gcflags="all=-N -l" -o /src/app


ADD entrypoint.sh /tmp/entrypoint.sh
RUN chmod 770 /tmp/entrypoint.sh

ENTRYPOINT ["/tmp/entrypoint.sh"]