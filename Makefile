
MKL_RED?=	\033[031m
MKL_GREEN?=	\033[032m
MKL_YELLOW?=	\033[033m
MKL_BLUE?=	\033[034m
MKL_CLR_RESET?=	\033[0m

BIN=      hashcode2017
prefix?=  /usr/local
bindir?=	$(prefix)/bin

build: vendor
	@printf "$(MKL_YELLOW)[BUILD]$(MKL_CLR_RESET)    Building project\n"
	@go build -ldflags "-X main.githash=`git rev-parse HEAD` -X main.version=`git describe --tags --always --dirty=-dev`" -o $(BIN) ./cmd
	@printf "$(MKL_YELLOW)[BUILD]$(MKL_CLR_RESET)    $(BIN) created\n"

install: build
	@printf "$(MKL_YELLOW)[INSTALL]$(MKL_CLR_RESET)  Installing $(BIN) to $(bindir)\n"
	@install $(BIN) $(bindir)
	@printf "$(MKL_YELLOW)[INSTALL]$(MKL_CLR_RESET)  Installed\n"

uninstall:
	@printf "$(MKL_RED)[UNINSTALL]$(MKL_CLR_RESET)  Remove $(BIN) from $(bindir)\n"
	@rm $(bindir)/$(BIN)
	@printf "$(MKL_RED)[UNINSTALL]$(MKL_CLR_RESET)  Removed\n"

tests:
	@printf "$(MKL_GREEN)[TESTING]$(MKL_CLR_RESET)  Running tests\n"
	@go test -race -v ./internal/...

coverage:
	@printf "$(MKL_BLUE)[COVERAGE]$(MKL_CLR_RESET) Computing coverage\n"
	@go test ./internal/... -covermode=count -coverprofile=coverage.out
	@go tool cover -func coverage.out

GLIDE := $(shell command -v glide 2> /dev/null)
vendor:
ifndef GLIDE
	$(error glide is not installed)
endif
	@printf "$(MKL_BLUE)[DEPS]$(MKL_CLR_RESET)  Resolving dependencies\n"
	@glide update

clean:
	rm -f $(BIN)
	rm -f coverage.out

vendor-clean:
	rm -rf vendor/

all: build tests coverage
