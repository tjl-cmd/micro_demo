package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/jinzhu/gorm"
	"user/common"
	"user/domain/repository"
	"user/handler"

	pb "user/proto/user"

	_ "github.com/go-sql-driver/mysql"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	service2 "user/domain/service"
)

var (
	service = "go.micro.service.user"
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
	// 初始化服务
	srv.Init()
	// 创建数据库连接
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

	rp := repository.NewUserRepository(db)
	rp.InitTable()
	UserDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	// Register handler
	pb.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: UserDataService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
