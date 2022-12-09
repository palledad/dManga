FROM ubuntu:22.04

RUN apt-get update
RUN apt-get install -y wget curl gnupg2

RUN mkdir /scripts \
    && wget 'https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh' -P /scripts  \
    && chmod 755 /scripts/wait-for-it.sh

WORKDIR /work

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate /usr/bin/migrate