module protoc-gen-tuyoo

go 1.15

require (
	github.com/google/btree v1.0.0 // indirect
	github.com/hashicorp/consul/api v1.2.0 // indirect
	github.com/liangdas/mqant v1.4.12
	github.com/prometheus/common v0.6.0
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	google.golang.org/protobuf v1.26.0
	sigs.k8s.io/yaml v1.1.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
