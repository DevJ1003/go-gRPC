package main

import (
	"log"
	"net"

	pb "github.com/devj1003/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:5051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %s\n", addr)

	// ==============================================================
	opts := []grpc.ServerOption{}
	tls := true // can change to false if needed
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.cpm"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}
	// ==============================================================

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
