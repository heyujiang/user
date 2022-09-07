BUILD_DATE := `date +%Y-%m-%dT%H:%M:%S%z`
BUILD_MANCHINE := `uname -mnsr`
GIT_VERSION := `git --no-pager describe --tags --dirty --always`
GIT_BRANCH := `git symbolic-ref --short HEAD`
VERSIONFILE := pkg/setting/version.go

LDFLAGS := -s -w
ifeq ($(shell uname), Linux)
ARM64_CROSS_COMPILE=aarch64-linux-gnu-
ARM_CROSS_COMPILE=arm-linux-gnueabihf-
endif

ifeq ($(shell uname), Darwin)
ARM64_CROSS_COMPILE=aarch64-unknown-linux-gnueabi-
ARM_CROSS_COMPILE=arm-unknown-linux-gnueabihf-
endif

all: version dependencies
	env CGO_ENABLED=1 go build

386: version dependencies
	env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build

amd64: version dependencies
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

arm: version dependencies
	env CC=${ARM_CROSS_COMPILE}gcc CXX=${ARM_CROSS_COMPILE}g++ CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -ldflags "$(LDFLAGS)"

aarch64: version dependencies
	#AARCH64_HOME is the absolute path for aarch64 compiler
	env CC=${ARM64_CROSS_COMPILE}gcc CXX=aarch64-linux-gnu-g++ CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CGO_LDFLAGS="-L${AARCH64_HOME}/lib" CGO_CFLAGS="-I${AARCH64_HOME}/include" go build


web:


pack:


version:
	rm -f $(VERSIONFILE)
	@echo "package setting" > $(VERSIONFILE)
	@echo "const (" >> $(VERSIONFILE)
	@echo "  GitVersion = \"$(GIT_VERSION)\"" >> $(VERSIONFILE)
	@echo "  GitBranch = \"$(GIT_BRANCH)\"" >> $(VERSIONFILE)
	@echo "  AppBuildTime = \"$(BUILD_DATE)\"" >> $(VERSIONFILE)
	@echo "  AppBuilder = \"$(USER)@'$(BUILD_MANCHINE)'\"" >> $(VERSIONFILE)
	@echo ")" >> $(VERSIONFILE)
	export GO111MODULE=on

dependencies:
	go generate -x ./...

clean:
	rm -rf *.db *.log FTU*
	find ./ -name *_string.go |xargs rm -rf
