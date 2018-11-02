.PHONY: lint
lint: fmt
	which golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $$GOPATH/bin v1.10
	golangci-lint run

# tests runs all test except those found in ./vendor
.PHONY: tests
tests: fmt 
	@echo "tests:"
	${GOPATH}/bin/richgo test ./...

# tests runs all test except those found in ./vendor
.PHONY: deltav
deltav: fmt 
	@echo "deltav:"
	${GOPATH}/bin/richgo test ./deltav/...


.PHONY: fmt
fmt:
	@echo "fmt:"
	scripts/fmt

.PHONY: server
server:	
	@echo "starting server (ctrl-C to exit)"
	go run deltav/mastercontrol/main/main.go

.PHONY: client
client:
	@echo "running client"
	go run deltav/shuttle/main/main.go

.PHONY: proto
proto:
	@echo "worldmodel:"
	protoc -I deltav/model \
	deltav/model/position.proto \
	deltav/model/vessel.proto \
	deltav/model/worldmodel.proto \
	--proto_path=. \
	--go_out=plugins=grpc:deltav/model/gomodel
