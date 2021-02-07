# Compile stage
FROM golang:latest AS build-env
# Build Delve
RUN go get github.com/go-delve/delve/cmd/dlv
ADD ./src /src
WORKDIR /src
RUN go build -gcflags="all=-N -l" -o /server
# Final stage
FROM debian:buster
EXPOSE 8000 40000
WORKDIR /
COPY --from=build-env /go/bin/dlv /
COPY --from=build-env /server /
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]
