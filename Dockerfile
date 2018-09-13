FROM golang

WORKDIR /go/src/github.com/Darkren/WeatherService

ADD . .

CMD ["go", "test", "./..."]