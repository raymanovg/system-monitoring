FROM ubuntu:18.04

LABEL MAINTAINERS="raymanovg@gmail.com"

RUN apt-get update \
    && apt-get install -y sysstat \
    && apt install -y iproute2

COPY bin/linux/amd64/monitoring /opt/monitoring/app

CMD /opt/monitoring/app