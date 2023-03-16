.PHONY: dist/server
dist/server: cmd/server/main.go
	go build -o $@ $<

.PHONY: init
init: weaver

.PHONY: weaver
weaver:
	go install github.com/ServiceWeaver/weaver/cmd/weaver

.PHONY: generate
generate: clean go-generate

.PHONY: go-generate
go-generate: # generate in the order of dependencies
	go generate ./weaverx/...
	go generate ./errorx/...
	go generate ./domain/...
	go generate ./persist/...
	go generate ./app/...
	go generate ./service/...

.PHONY: seed
seed:
	bin/seed.sh

.PHONY: clean
clean:
	find . -type f -name "*_generated.go" -delete
	find . -type f -name "weaver_gen.go" -delete

.PHONY: test
test:
	go test -v ./cmd/server
