package main

import (
	"helloworld/handler"
	pb "helloworld/proto"

	service "github.com/asim/go-micro/v3"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	//srv := service.New(
	//	service.Name("helloworld"),
	//	service.Version("latest"),
	//)

	srv := service.NewService(
		service.Name("helloworld"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
