APP:=ab_log_grapher
PROJECT?=ab_log_inserter
RELEASE?=0.0.1
#COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CGO_ENABLED := CGO_ENABLED=0

# File name
DIR:= bin

# Environment variables for build
ENV_OSX64 := GOOS=darwin GOARCH=amd64
ENV_OSX32 := GOOS=darwin GOARCH=386
ENV_LIN64 := GOOS=linux GOARCH=amd64
ENV_LIN32 := GOOS=linux GOARCH=386
ENV_WIN64 := GOOS=windows GOARCH=amd64
ENV_WIN32 := GOOS=windows GOARCH=386
ENV_ARM := GOOS=linux GOARCH=arm
# Build command
CMD_BUILD := go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o


windows:
	@$(CGO_ENABLED) $(ENV_WIN64) $(CMD_BUILD) $(DIR)/$(APP).exe
#	@$(CGO_ENABLED) $(ENV_WIN32) $(CMD_BUILD) $(DIR)/$(APP)win32.exe
	@echo "Windows complete."
	
	
linux:
	@$(CGO_ENABLED) $(ENV_LIN64) $(CMD_BUILD) $(DIR)/$(APP)
#	@$(CGO_ENABLED) $(ENV_LIN32) $(CMD_BUILD) $(DIR)/$(APP)lin32.bin
	@echo "Linux complete."
rasp:
	@$(CGO_ENABLED) $(ENV_ARM) $(CMD_BUILD) $(DIR)/$(APP)
clean:
			rm -r $(DIR)
