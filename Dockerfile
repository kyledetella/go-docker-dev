FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/kyledetella/go-docker-dev/app
WORKDIR /go/src/github.com/kyledetella/go-docker-dev/app

RUN go get ./
RUN go build

CMD go get github.com/pilu/fresh && fresh

EXPOSE 8080
