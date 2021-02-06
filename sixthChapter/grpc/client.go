package main

import (
	"context"
	"log"

	pb "github.com/sarkan9/moneyTransactionProto"

	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewMoneyTransactionClient(conn)
	from := "sarkan"
	to := "arash"
	var amount float32
	amount = 210.5
	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{
		From:   from,
		To:     to,
		Amount: amount,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The Request, Respones: %t", r.Confirmation)
}
