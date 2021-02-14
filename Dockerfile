FROM golang:1.15

RUN go get github.com/go-delve/delve/cmd/dlv

ADD ./src /src
WORKDIR /src

RUN go build -gcflags="all=-N -l" -o /src/app

EXPOSE 8000 40000

CMD ["dlv", "--listen=:40000", "--log", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/src/app"]