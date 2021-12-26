module product

go 1.16

require (
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.0.0-20211224114320-e0de23c54615
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211224114320-e0de23c54615
	github.com/jinzhu/gorm v1.9.16
	go-micro.dev/v4 v4.2.1
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
