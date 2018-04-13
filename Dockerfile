FROM golang:1.10-stretch

WORKDIR /go/src/github.com/gulien/orbit

# Installs lint dependencies.
RUN go get -u gopkg.in/alecthomas/gometalinter.v2 &&\
    gometalinter.v2 --install

# Installs dep for our tests.
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ENV SHELL="/bin/sh"

# Copies our Go source.
COPY . .

# Installs project dependencies.
RUN go get -d -v ./...

ENTRYPOINT [".ci/docker-entrypoint.sh"]