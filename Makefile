RWILDCARD=$(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call RWILDCARD,$d/,$2))
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(call RWILDCARD,internal/,*.proto)
API_PROTO_FILES=$(call RWILDCARD,api/,*.proto)
ERROR_PROTO_FILES=$(call RWILDCARD,api/,error*.proto)

.PHONY: init
# init env
init:
	go get -u github.com/go-kratos/kratos/cmd/kratos/v2
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/gogo/protobuf/gogoproto
	go get -u github.com/gogo/protobuf/protoc-gen-gogoslick



.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
                   --proto_path=./third_party \
                   --go_out=paths=source_relative:. \
                   --go-errors_out=paths=source_relative:. \
                   $(ERROR_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	@$(foreach proto_file,$(API_PROTO_FILES), \
		echo $(proto_file); \
        protoc  --proto_path=. \
                --proto_path=./third_party \
                --gogoslick_out=paths=source_relative:. \
                --go-xhttp_out=paths=source_relative:. \
                --go-grpc_out=paths=source_relative:. \
        $(proto_file); \
	)\
 	echo 'generate API proto file done!'

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
