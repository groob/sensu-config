FROM golang
WORKDIR /usr/src/go/src
RUN CGO_ENABLED=0 ./make.bash
RUN mkdir -p /go/src/sensu-config
COPY main.go /go/src/sensu-config/main.go
WORKDIR /go/src/sensu-config

