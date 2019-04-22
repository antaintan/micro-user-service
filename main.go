package main

import (
	"huochain/gomicro-demo/user-service/pb"
	"huochain/gomicro-demo/user-service/service"
	"log"

	"github.com/micro/go-config"
	"github.com/micro/go-micro"
)

func main() {
	config.LoadFile("app.yaml")
	conf := config.Map()

	log.Print(conf["server.name"])

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		// micro.RegisterTTL(time.Second*30),
		// micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	srv.Init()

	// Register Handlers
	pb.RegisterUserServiceHandler(srv.Server(), new(service.UserService))

	// Run server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
