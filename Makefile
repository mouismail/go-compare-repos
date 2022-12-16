# define the binary name
BINARY=go-action-runner

# define the build directory
BUILD_DIR=build

# define the source files
SRCS= $(wildcard *.go)

# define the build flags
BUILD_FLAGS=-ldflags "-X main.Version=$(shell git describe --tags --dirty --always)"

# build the application
build: $(SRCS)
    mkdir -p $(BUILD_DIR)
    env GH_ORG_NAME=$(GH_ORG_NAME) GL_PROJECT_ID=$(GL_PROJECT_ID) GITLAB_ACCESS_TOKEN=$(GITLAB_ACCESS_TOKEN) GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY) $(SRCS)

# run the application
run:
    env GH_ORG_NAME=$(GH_ORG_NAME) GL_PROJECT_ID=$(GL_PROJECT_ID) GITLAB_ACCESS_TOKEN=$(GITLAB_ACCESS_TOKEN) GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) $(BUILD_DIR)/$(BINARY)

# run the tests
test:
    go test -v ./...

# clean up the build directory
clean:
    rm -rf $(BUILD_DIR)

# default rule to build the application
default: build
