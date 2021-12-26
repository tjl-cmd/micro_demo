package main

import (
	"cart/handler"
	pb "cart/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "cart"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterCartHandler(srv.Server(), new(handler.Cart))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
