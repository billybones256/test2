package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"task2/pkg/api"
	"task2/pkg/pyramid"
)

func main() {
	server := grpc.NewServer()
	microService := &pyramid.GRPCServer{}
	api.RegisterPyramidServer(server, microService)
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
