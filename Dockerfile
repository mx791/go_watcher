FROM golang:1.12-alpine AS build_base

COPY . /usr/src/app

WORKDIR /usr/src/app
RUN go build
RUN ls -la

FROM alpine
COPY --from=build_base /usr/src/app /usr/src/app

WORKDIR /usr/src/app

CMD ./app