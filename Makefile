include .env
SHELL = bash
.PHONY := default init test build release env env-push
.DEFAULT_GOAL = default

PROJECT := $(shell basename $$PWD)
GOTEST ?= go test

default:
	@ mmake help

# init project
init: 
	@ go mod vendor

# run tests
test:
	${GOTEST} ./...

# create release VERSION on github
#
# VERSION should being with a v and be in SemVer format.
release:
	$(eval VERSION=$(filter-out $@, $(MAKECMDGOALS)))
	$(if ${VERSION},@true,$(error "VERSION is required"))
	git commit --allow-empty -am ${VERSION}
	git push
	hub release create -m ${VERSION} -e ${VERSION}

docs:
	godoc -http=:6060 -goroot=${HOME}/code/gb	

%:
	@ true

