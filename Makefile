CI_IMAGE_NAME = gogfapi-ci
CONTAINER_CMD := docker
CONTAINER_NAME := gluster-test
CONTAINER_CONFIG_DIR := testing/containers

GFAPI_APIS := enhanced
GFAPI_APIS_TEST_TAGS :=
GLUSTER_VERSION := latest

ifeq ($(GFAPI_APIS),legacy)
	GLUSTER_VERSION = gluster-4.1
	GFAPI_APIS_TEST_TAGS = -tags glusterfs_legacy_api
endif

# the name of the image plus ceph version as tag
CI_IMAGE_TAG=$(CI_IMAGE_NAME):$(GLUSTER_VERSION)

test:
	go test -v ./...

.PHONY: ci-image
ci-image:
	$(CONTAINER_CMD) build --build-arg GLUSTER_VERSION=$(GLUSTER_VERSION) -t $(CI_IMAGE_TAG) -f $(CONTAINER_CONFIG_DIR)/Dockerfile .

.PHONY: setup-test-container
setup-test-container:
	CI_IMAGE_TAG=${CI_IMAGE_TAG} CONTAINER_NAME=${CONTAINER_NAME} ./testing/test_setup.sh

test-container:
	$(CONTAINER_CMD) exec ${CONTAINER_NAME} go test ${GFAPI_APIS_TEST_TAGS} -v ./...