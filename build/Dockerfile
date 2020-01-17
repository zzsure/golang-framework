FROM ubuntu:xenial
RUN apt-get update
RUN apt-get install tzdata
RUN echo "Asia/Shanghai" > /etc/timezone
RUN rm -f /etc/localtime
RUN dpkg-reconfigure -f noninteractive tzdata
RUN apt-get install -y ca-certificates
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai/etc/localtime
#RUN ln -fs /usr/share/zoneinfo/US/Pacific-New /etc/localtime && dpkg-reconfigure -f noninteractive tzdata
WORKDIR /root/deploy/golang-golang-framework

RUN mkdir -p /root/deploy/golang-golang-framework
COPY config.toml /root/deploy/golang-golang-framework/config.toml
COPY golang-golang-framework /root/deploy/golang-golang-framework/
RUN chmod +x /root/deploy/golang-golang-framework/golang-golang-framework

ENTRYPOINT ["/root/deploy/golang-golang-framework/golang-golang-framework"]
