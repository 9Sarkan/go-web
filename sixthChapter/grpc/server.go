package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	pb "github.com/sarkan9/moneyTransactionProto"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct{}

// MakeTransaction method
func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Println("request for transaction...")
	log.Printf("amount: %f, from: %s, to: %s\n", in.Amount, in.From, in.To)
	return &pb.TransactionResponse{Confirmation: true}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
