FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/EricWang-Git/bike
RUN go install github.com/EricWang-Git/bike

EXPOSE 80
CMD ["/gopath/app/bin/echo"]
