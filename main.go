package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/refs/exposer/pkg/exposersvc"
	exposer "github.com/refs/exposer/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.exposer"), // TODO fetch this from the imported service
	)

	service.Init()
	exposer.RegisterExposerHandler(service.Server(), &exposersvc.ExposerHandler{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
