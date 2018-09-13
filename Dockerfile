FROM golang

WORKDIR /go/src/WeatherService

ADD . .

RUN echo $GOPATH
RUN echo $GOROOT

CMD ["go", "test", "./..."]