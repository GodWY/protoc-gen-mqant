GOPATHS=`go env GOPATH`
init:
	@echo "init mqant tools"

windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "protoc-gen-mqant" main.go
	@cp  protoc-gen-mqant $(GOPATHS)/bin
	@rm  protoc-gen-mqant
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o "protoc-gen-mqant" main.go
	@cp  protoc-gen-mqant $(GOPATHS)/bin
	@rm  protoc-gen-mqant
mac:
	go build -o "protoc-gen-mqant" main.go
	@cp  protoc-gen-mqant $(GOPATHS)/bin
	@rm  protoc-gen-mqant

install:
	@echo $(GOPATHS)
