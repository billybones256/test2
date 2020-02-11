package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	"task2/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	h, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	connection, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewPyramidClient(connection)
	res, err := client.BuildPyramid(context.Background(), &api.PyramidRequest{H: int32(h)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetAnswer())
}
