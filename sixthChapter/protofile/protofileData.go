package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/sarkan9/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1,
		Name:  "sarkan",
		Email: "sarkan@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Phone: "09184721585", Type: pb.Person_Mobile},
		},
	}
	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("original data: ", p)
	fmt.Println("marshals proto data: ", body)
	fmt.Println("unmarshal struct: ", p1)
}
