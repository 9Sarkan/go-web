package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/reflection"

	pb "github.com/sarkan9/streamProto"
	"google.golang.org/grpc"
)

type server struct{}

const port = ":50051"
const noOfSteps = 3

// MakeTransaction function
func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	log.Println("got new request...")
	log.Printf("From: %s\tTo: %s\tAmount: %v\n", in.From, in.To, in.Amount)
	for i := 0; i <= noOfSteps; i++ {
		time.Sleep(2 * time.Second)
		if err := stream.Send(&pb.TransactionResponse{
			Status:      "Good",
			Step:        int32(i),
			Description: fmt.Sprintf("Description of step: %v", int32(i)),
		}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}
	log.Println("request responsed successfully!")
	return nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
