FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ADD . /gopath/app/

EXPOSE 80
CMD ["/gopath/app/bin/echo"]