# WeatherService
Starts on the port specified in config. Endpoints:
1. Start search
/weather/search [POST] 
{
    "lat": 12.56,
    "lon": 12.78
}

2. Get results
/weather/result/{requestId} [GET]


Data storage can be changed with new implementations of repository interfaces
External weather info source can be changed with new implementation of weather service interface
More than one worker can be launched in order to make processing faster. Just create and run more of them -
everything should work just fine


Building app:
1. dep ensure
2. go build

Starting locally:
1. Create database in PgSQL, also change the dbName in config (db section)
2. Run dbinit/dbinit.sql on this database
3. Change db config (json db section) to fit the current db setup
4. ./weatherservice

Running in container:
1. cd dbinit (assuming you are starting in the project dir)
2. sudo docker build -t pgsql/weatherservice .
3. sudo docker run -itd --name pgsql pgsql/weatherservice
4. cd ../
5. Change json.config to target port 5432, password specified in Dockerfile for db (1234 by default), host pgsql:
    "db": {
        "host": "pgsql",
        "port": 5432,
        "user": "postgres",
        "password": "1234",
        "dbName": "weatherservice",
        "sslmode": "disable"
    }
6. go build
7. sudo docker build -t golang/weatherservice .
8. sudo docker run -itd --name weatherservice --link pgsql:pgsql/weatherservice -p 8080:8080 golang/weatherservice