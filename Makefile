

# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)


# --- Tooling & Variables ----------------------------------------------------------------
include ./misc/make/help.Makefile

# ~~~ Development Environment ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

up: run             ## Startup / Spinup Docker Compose and air
down: docker-stop               ## Stop Docker
destroy: docker-teardown clean  ## Teardown (removes volumes, tmp files, etc...)

run: ## Bootstrap Environment (with a Docker-Compose help).
	@ go test ./...
	@ gosec ./...
	@ docker-compose up -d --build

docker-stop:
	@ docker-compose down

docker-teardown:
	@ docker-compose down --remove-orphans -v

# ~~~ Code Actions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

build: ## Builds binary
	@ printf "Building aplication... "
	@ go build \
		-trimpath  \
		-o dating-apps-api \
		./app/
	@ echo "done"

TESTS_ARGS := --format testname --jsonfile gotestsum.json.out
TESTS_ARGS += --max-fails 2
TESTS_ARGS += -- ./...
TESTS_ARGS += -test.parallel 2
TESTS_ARGS += -test.count    1
TESTS_ARGS += -test.failfast
TESTS_ARGS += -test.coverprofile   coverage.out
TESTS_ARGS += -test.timeout        5s
TESTS_ARGS += -race

run-tests: 
	@go test ./...

# ~~~ Docker Build ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.ONESHELL:
image-build:
	@ echo "Docker Build"
	@ DOCKER_BUILDKIT=0 docker build \
		--file Dockerfile \
		--tag go-clean-arch \
			.


clean: clean-artifacts clean-docker

clean-artifacts: ## Removes Artifacts (*.out)
	@printf "Cleanning artifacts... "
	@rm -f *.out
	@echo "done."


clean-docker: ## Removes dangling docker images
	@ docker image prune -f
