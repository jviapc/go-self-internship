PRJ_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))/src
IMAGE_NAME := gotest

install:
	docker build -t ${IMAGE_NAME} --build-arg HOST_UID=$(shell id -u) ./

test:
	docker run --rm -it --tty -v ${PRJ_DIR}:/app composer test

clean:
	git clean -fdx -e .idea

terminal:
	docker run --rm -it --tty -v ${PRJ_DIR}:/app ${IMAGE_NAME} sh

