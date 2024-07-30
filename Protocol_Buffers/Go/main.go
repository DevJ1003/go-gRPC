package main

import (
	"fmt"

	pb "github.com/devj1003/proto-go-course/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 42, Name: "My Name"},
		MultipleDummies: []*pb.Dummy{
			{Id: 43, Name: "Second Name"},
			{Id: 44, Name: "Third Name"},
		},
	}
}

func main() {
	// fmt.Println(doSimple())
	fmt.Println(doComplex())
}
