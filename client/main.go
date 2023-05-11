package main

import (
	"context"
	api_transfrom "github.com/aesoper101/kitextest/client/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	ctx := context.Background()
	req := &api_transfrom.TransferRequest{
		Amount:   10,
		FromUser: "user1",
		ToUser:   "user2",
	}

	dial, err := grpc.Dial(":8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("dial err:%s", err.Error())
		return
	}

	client := api_transfrom.NewTransferServiceClient(dial)

	for i := 0; i < 10; i++ {
		m, err := client.TransIn(ctx, req)
		if err != nil {
			log.Printf("call err:%s", err.Error())
			return
		}

		log.Printf("call resp:%s", m.Reason)
	}

	_ = dial.Close()

}
