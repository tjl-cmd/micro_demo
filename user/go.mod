module user

go 1.16

require (
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.0.0-20211220083148-8e52761edb49
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211220083148-8e52761edb49
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	go-micro.dev/v4 v4.2.1
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
