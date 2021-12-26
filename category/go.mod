module category

go 1.16

require (
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.0.0-20211210113221-37de747d195c
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211210113221-37de747d195c
	github.com/go-sql-driver/mysql v1.5.0
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.3.0 // indirect
	go-micro.dev/v4 v4.2.1
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
