FROM postgres

WORKDIR /docker-entrypoint-initdb.d

ENV POSTGRES_DB weatherservice

ADD dbinit.sql .

CMD ["postgres"]