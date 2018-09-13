FROM golang

ADD . /go/src/WeatherService

WORKDIR /go/src/WeatherService

RUN ls -la

CMD ["go", "test", "./..."]