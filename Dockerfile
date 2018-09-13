FROM golang

ADD . /go/src/WeatherService

WORKDIR /go/src/WeatherService

CMD ["go", "test", "./..."]