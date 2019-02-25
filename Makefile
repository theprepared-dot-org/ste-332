.PHONY: wiki build

wiki: tools
	git submodule update --init --recursive
	hugo -s wiki

serve-wiki:
	hugo -s wiki server

tools: $GOPATH/bin/hugo $GOPATH/bin/protoc-gen-go

$GOPATH/bin/hugo:
	go install github.com/gohugoio/hugo

$GOPATH/bin/protoc-gen-go:
	go install github.com/golang/protobuf/protoc-gen-go

clean:
	rm -rf wiki/resources
	rm -rf wiki/public

rpc:
	go get google.golang.org/grpc
	protoc -I api/shopmanager/ api/shopmanager/shop_manager.proto --go_out=plugins=grpc:api/shopmanager

build:
	go build -o build/out/shop-manager github.com/theprepared-dot-org/ste332/cmd/shopmanager

test:
	go test -v -cover -short ./...

test-full:
	go test -v -cover ./...

run:
	go run cmd/shopmanager/main.go
