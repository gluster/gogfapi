#!/bin/bash

docker run -d --name gluster-test \
  -p 24007:24007/tcp -p 24008:24008/tcp -p 24007:24007/udp -p 24008:24008/udp \
  -p 49152:49152/tcp -p 49152:49152/udp -p 49153:49153/tcp -p 49154:49154/tcp \
  -v /sys/fs/cgroup:/sys/fs/cgroup:ro --privileged=true \
  gluster/gluster-centos:gluster4u0_centos7

# wait for glusterd service to be active
while ! docker exec gluster-test systemctl is-active glusterd; do
  sleep 1
done

docker exec gluster-test /bin/sh -c "echo '127.0.1.1 $(hostname)' >> /etc/hosts"
docker exec gluster-test gluster volume create test $(hostname):/srv force
docker exec gluster-test gluster volume set test storage.owner-uid $(id -u)
docker exec gluster-test gluster volume start test
