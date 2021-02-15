FROM golang:1.15

RUN go get github.com/go-delve/delve/cmd/dlv

ADD ./src /src
WORKDIR /src

RUN go mod download

EXPOSE 8000 40000

ADD ./entrypoint.sh /tmp/entrypoint.sh
RUN chmod 770 /tmp/entrypoint.sh

ENTRYPOINT [ "/tmp/entrypoint.sh" ]
