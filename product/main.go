package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/jinzhu/gorm"
	"go-micro.dev/v4/registry"
	"product/common"
	"product/domain/repository"
	service2 "product/domain/service"
	"product/handler"
	pb "product/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "go.micro.product.service"
	version = "latest"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulRegistry),
	)
	// 获取mysql配置,路径中不带前缀
	sql := common.GetMysqlFromConsul(consulConfig, "mysql")
	if sql.User == "" || sql.Host == "" || sql.Port == 0 || sql.Pwd == "" || sql.Database == "" {
		log.Error("初始化配置失败")
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", sql.User, sql.Pwd, sql.Host, sql.Port, sql.Database)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)
	srv.Init()
	rp := repository.NewProductRepository(db)
	err = rp.InitTable()
	if err != nil {
		log.Error("初始化数据库表失败")
	}
	productDataService := service2.NewProductDataService(rp)
	// Register handler
	err = pb.RegisterProductHandler(srv.Server(), &handler.Product{ProductDataService: productDataService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
