FROM postgres

WORKDIR /docker-entrypoint-initdb.d

ENV POSTGRES_DB weatherservice
ENV POSTGRES_PASSWORD 1234

ADD dbinit.sql .

CMD ["postgres"]