FROM debian:jessie
MAINTAINER Roman Safronov <electroma@gmail.com>

ENV DEBIAN_FRONTEND noninteractive

# Installing dependencies
RUN apt-get update && apt-get install -y libldap-2.4-2 libkrb5-3 libsasl2-2 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY go-ad-man /
COPY views /views
ENTRYPOINT ["/go-ad-man"]
