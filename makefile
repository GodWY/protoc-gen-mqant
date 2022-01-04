GOPATHS=`go env GOPATH`
init:
	@echo "init mqant tools"

windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "protoc-gen-mqant" main.go
	@cp  mqant $(GOPATHS)/bin
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o "protoc-gen-mqant" main.go
	@cp  mqant $(GOPATHS)/bin
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -o "protoc-gen-mqant" main.go
	@cp  mqant $(GOPATHS)/bin

install:
	@echo $(GOPATHS)
