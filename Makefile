COMMANDS=flux
BINARIES=$(addprefix bin/,$(COMMANDS))
PKG=github.com/ehotinger/flux
PACKAGES=$(shell go list ./... | grep -v /vendor/)
GOPATHS=$(shell echo ${GOPATH} | tr ":" "\n" | tr ";" "\n")
GO_GCFLAGS=$(shell				\
	set -- ${GOPATHS};			\
	echo "-gcflags=-trimpath=$${1}/src";	\
	)
GO_BUILD_FLAGS=
GO_LDFLAGS=-ldflags '-s -w -X $(PKG)/version.Version=$(VERSION) -X $(PKG)/version.Revision=$(REVISION) -X $(PKG)/version.Package=$(PKG) $(EXTRA_LDFLAGS)'

all: binaries

FORCE:

binaries: $(BINARIES) ## builds bin/ binaries

bin/%: cmd/% FORCE
	@echo "+ $@${BINARY_SUFFIX}"
	@go build ${GO_GCFLAGS} ${GO_BUILD_FLAGS} -o $@${BINARY_SUFFIX} ${GO_LDFLAGS} ${GO_TAGS}  ./$<

clean: ## cleans remnants of binaries
	@echo "+ $@"
	@rm -f $(BINARIES)

help: ## print this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort
