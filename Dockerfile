FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/goufu1989/echo
RUN go install github.com/goufu1989/echo

EXPOSE 80
CMD ["/gopath/app/bin/echo"]