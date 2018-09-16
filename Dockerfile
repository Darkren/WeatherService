FROM golang

WORKDIR /go/src/github.com/Darkren/weatherservice

ADD . .

RUN go test -v -race ./...

EXPOSE 8080

CMD ["./weatherservice"]

#sudo docker run --name weatherservice --link pgsql:postgres -p 8080:8080 golang/app
#sudo docker build -t pgsql/weatherservice .
#sudo docker run --name pgsql -p 5433:5432 -e POSTGRES_PASSWORD=1234 pgsql/weatherservice