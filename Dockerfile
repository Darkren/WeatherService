FROM golang

RUN mkdir /app

COPY . /app/

WORKDIR /app

CMD ["go", "test", "./..."]