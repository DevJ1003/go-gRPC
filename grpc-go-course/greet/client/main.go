package main

import (
	"log"

	pb "github.com/devj1003/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:5051"

func main() {

	// ==================================================
	tls := true // can be changed to false if needed
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	// ==================================================

	// conn, err := grpc.Dial(addr)
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	doGreet(c) // also used for SSL security
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)

	// doGreetWithDeadline(c, 5*time.Second)
	// doGreetWithDeadline(c, 1*time.Second)
}
