package main

import (
	"context"
	"fmt"
	api_transfrom "github.com/aesoper101/kitextest/streamclient/api"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	ctx := context.Background()

	dial, err := grpc.Dial(":8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("dial err:%s", err.Error())
		return
	}

	transferServiceClient := api_transfrom.NewTransferServiceClient(dial)
	tt := time.After(time.Second * 10)

	for {
		select {
		case <-tt:
			return
		default:
			req := &api_transfrom.TransferRequest{
				Amount:   10,
				FromUser: "user1",
				ToUser:   "user2",
			}
			stream, err := transferServiceClient.TransIn(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = stream.Send(req)
			if err != nil {
				fmt.Println(err)
				return
			}

			resp := &api_transfrom.Message{}
			err = stream.RecvMsg(resp)
			if err != nil {
				fmt.Println("recv err:", err)
				return
			}

			fmt.Println("recv resp:", resp.Reason)
		}
	}

	_ = dial.Close()
}
