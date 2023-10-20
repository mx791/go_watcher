FROM golang:1.12-alpine AS build_base

COPY . /usr/src/app

WORKDIR /usr/src/app
RUN go build
RUN ls -la

FROM alpine
RUN apk update && apk upgrade && apk add --no-cache bash git openssh

WORKDIR /usr/src/app

ENV TARGET_DIR_PATH "/git_path"
ENV GIT_URL ""
ENV PERIOD 60
ENV POST_UPDATE ""

COPY --from=build_base /usr/src/app /usr/src/app

CMD ./app