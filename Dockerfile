#  解决 docker
FROM golang:latest as builder
# FROM fagongzi/golang:1.6.1 as builder
COPY . /go/src/github.com/matrixorigin/talent-challenge/matrixbase/distributed
WORKDIR /go/src/github.com/matrixorigin/talent-challenge/matrixbase/distributed

# 解决 etcd带来的 docker 编译错误
RUN mkdir -p /home/luluyuzhi/go/src/github.com/etcd-io/
RUN cp -r ./extra/* /home/luluyuzhi/go/src/github.com/etcd-io/
RUN ls -l /home/luluyuzhi/go/src/github.com/etcd-io/

RUN make

FROM alpine:latest

RUN mkdir -p /apps/matrixbase/distributed/bin
RUN mkdir -p /apps/matrixbase/distributed/logs

ARG APP_ROOT=/apps/matrixbase/distributed
ENV PATH=${APP_ROOT}/bin:$PATH

WORKDIR ${APP_ROOT}

COPY --from=builder /go/src/github.com/matrixorigin/talent-challenge/matrixbase/distributed/dist/server ${APP_ROOT}/bin/server

# Alpine Linux doesn't use pam, which means that there is no /etc/nsswitch.conf,
# but Golang relies on /etc/nsswitch.conf to check the order of DNS resolving
# (see https://github.com/golang/go/commit/9dee7771f561cf6aee081c0af6658cc81fac3918)
# To fix this we just create /etc/nsswitch.conf and add the following line:
# hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4

RUN echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf

ENTRYPOINT ["/apps/matrixbase/distributed/bin/server"]
