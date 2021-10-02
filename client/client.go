package main

import (
	"google.golang.org/grpc"
)

func main() {
	_, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
}
