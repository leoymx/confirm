FROM google/golang

# Build app
WORKDIR /gopath/src/app
ADD . /gopath/src/app/

RUN go get app
RUN go install app

EXPOSE 80
CMD ["/gopath/app/bin/app"]