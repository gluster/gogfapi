#!/bin/bash

docker run -d --name ${CONTAINER_NAME} \
  -v /sys/fs/cgroup:/sys/fs/cgroup:ro --privileged=true \
  -v $(pwd):/go/src/github.com/gluster/gogfapi \
  ${CI_IMAGE_TAG}

# wait for glusterd service to be active
while ! docker exec ${CONTAINER_NAME} systemctl is-active glusterd; do
  sleep 1
done

docker exec ${CONTAINER_NAME} /bin/sh -c "echo '127.0.1.1 $(hostname)' >> /etc/hosts"
docker exec ${CONTAINER_NAME} gluster volume create test $(hostname):/srv force
docker exec ${CONTAINER_NAME} gluster volume set test storage.owner-uid $(id -u)
docker exec ${CONTAINER_NAME} gluster volume start test


