package main

import (
	"category/common"
	"category/domain/repository"
	service2 "category/domain/service"
	"category/handler"
	pb "category/proto"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	service = "go.micro.service.category"
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
		micro.Address("127.0.0.1:8082"),
		// 添加consul，作为注册中心
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

	rp := repository.NewCategoryRepository(db)
	err = rp.InitTable()
	if err != nil {
		log.Error("初始化数据库表失败")
	}
	categoryDataService := service2.NewCategoryDataService(rp)
	// Register handler
	err = pb.RegisterCategoryHandler(srv.Server(), &handler.Category{CategoryDataService: categoryDataService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
