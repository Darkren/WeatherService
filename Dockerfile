FROM golang

WORKDIR /go/src/github.com/Darkren/weatherservice

ADD . .

CMD ["go", "test", "./..."]