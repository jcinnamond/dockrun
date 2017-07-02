FROM ruby:2.3.3

ADD . /app
WORKDIR /app

ENTRYPOINT "bash"
