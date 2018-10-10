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

.PHONY: common
common:
	@echo "common:"
	protoc -I deltav/common/ deltav/common/position.proto --go_out=plugins=grpc:deltav/common


.PHONY: mc
mc:
	@echo "mastercontrol:"
	protoc -I deltav/mastercontrol/ deltav/mastercontrol/worldmodel.proto --go_out=plugins=grpc:deltav/mastercontrol
