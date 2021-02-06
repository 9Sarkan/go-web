package main

import (
	"io"
	"log"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	pb "github.com/sarkan9/streamProto"
)

const (
	address = "localhost:50051"
)

// ReceiveStream function
func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("start listen to server stream...")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("status: %s -- step: %d -- Description: %s", response.Status, response.Step, response.Description)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewMoneyTransactionClient(conn)
	request := pb.TransactionRequest{
		From:   "sarkan",
		To:     "Arash",
		Amount: 123.12,
	}
	ReceiveStream(client, &request)
}
