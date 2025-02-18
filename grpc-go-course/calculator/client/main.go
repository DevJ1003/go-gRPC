package main

import (
	"log"

	pb "github.com/devj1003/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5051"

func main() {

	// conn, err := grpc.Dial(addr)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	// doSum(c)
	// doPrimes(c)
	// doAverage(c)
	// doMax(c)

	// doSqrt(c, 10)
	doSqrt(c, -2)
}
