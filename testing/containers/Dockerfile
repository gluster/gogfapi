ARG GLUSTER_VERSION
FROM quay.io/gluster/gluster-centos:${GLUSTER_VERSION:-latest}

RUN true && \
    yum clean all && \
    gv="$(rpm -q --queryformat '%{version}-%{release}' glusterfs-server)" && \
    yum install -y \
        git wget make gcc curl "glusterfs-api-devel-${gv}" && \
    true

ENV GOTAR=go1.14.9.linux-amd64.tar.gz
RUN true && \
    curl -o /tmp/${GOTAR} https://dl.google.com/go/${GOTAR} && \
    tar -x -C /opt/ -f /tmp/${GOTAR} && \
    rm -f /tmp/${GOTAR} && \
    true

ENV PATH="${PATH}:/opt/go/bin"
ENV GOROOT=/opt/go
ENV GO111MODULE=off
ENV GOPATH /go
WORKDIR /go/src/github.com/gluster/gogfapi
VOLUME /go/src/github.com/gluster/gogfapi