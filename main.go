package main

import (
	transfrom "github.com/aesoper101/kitextest/kitex_gen/transfrom/transferservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	svr := transfrom.NewServer(
		new(TransferServiceImpl),
		server.WithServiceAddr(addr),
		// 指定了以下两个选项，都没有用
		//server.WithCodec(protobuf.NewGRPCCodec()),
		//server.WithPayloadCodec(protobuf.NewProtobufCodec()),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
