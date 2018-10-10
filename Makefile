.PHONY: lint
lint: fmt
	which golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $$GOPATH/bin v1.10
	golangci-lint run

# tests runs all test except those found in ./vendor
.PHONY: tests
tests: fmt 
	@echo "tests:"
	${GOPATH}/bin/richgo test ./...


.PHONY: fmt
fmt:
	@echo "fmt:"
	scripts/fmt

.PHONY: proto
proto:
	@echo "worldmodel:"
	protoc -I deltav/protos \
	deltav/protos/position.proto \
	deltav/protos/worldmodel.proto \
	--proto_path=. \
	--go_out=plugins=grpc:deltav/protos
