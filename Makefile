.PHONY: build run test test-coverage test-show-coverage check-prereqs swagger protos bindata deps certs docker stop-docker hosts

build:
	@./_scripts/build.sh

run: check-prereqs certs docker deps protos bindata swagger
	@./_scripts/service.sh

tests:
	@./_scripts/tests_unit.sh

test-coverage:
	@./_scripts/tests_unit.sh total

test-show-coverage:
	@./_scripts/tests_coverage.sh
	@./_scripts/tests_coverage.sh service

# test-show-coverage-service:
# 	@./_scripts/tests_coverage.sh service

check-prereqs:
	@./_scripts/prereq-check.sh

swagger:
	@./_scripts/swagger.sh

protos:
	@./_scripts/protos.sh

bindata:
	@./_scripts/bindata.sh

certs:
	@./_scripts/certs.sh

hosts:
	@./_scripts/hosts.sh

deps:
	@echo "Verifying go mod dependencies"
# @go mod tidy
# @go mod vendor
# If it were a private repo
#	@GOPROXY=direct GOPRIVATE=github.com/stevepartridge/blogorama/* go mod tidy
#	@GOPROXY=direct GOPRIVATE=github.com/stevepartridge/blogorama/* go mod vendor

docker:
	@echo docker
	@# make sure we have 2
	@if [ $(shell docker ps -a --no-trunc --quiet --filter name=^/blog --filter status=running | wc -l) -eq 2 ]; then \
		echo Containers runnning.; \
	else \
		./_scripts/compose.sh; \
	fi

stop-docker:
	@echo Stopping Docker...
	@# see if we have 1 or more
	@if [ $(shell docker ps -a --no-trunc --quiet --filter name=^/blog --filter status=running | wc -l) -gt 0 ]; then \
		./_scripts/compose.sh stop; \
	else \
		echo "Nothing to do."; \
	fi

clean: stop-docker
	@echo clean volumes
	@./_scripts/clean.sh