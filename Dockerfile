FROM golang:1.16-alpine AS build

WORKDIR /usr/src/app

COPY . /usr/src/app/

CMD ["go", "run", "."]