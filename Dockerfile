
FROM ubuntu:latest

MAINTAINER Mickey Yawn <mickey.yawn@turner.com>

RUN apt-get update && apt-get install -y supervisor

RUN mkdir -p /var/log/supervisor

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

ADD backplane-poc /backplane-poc

ADD backplane /backplane

CMD ["/usr/bin/supervisord"]