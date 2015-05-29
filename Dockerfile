FROM google/golang
MAINTAINER wangwei

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/leoymx/confirm
RUN go install github.com/leoymx/confirm

EXPOSE 80
CMD ["/gopath/app/bin/confirm"]
