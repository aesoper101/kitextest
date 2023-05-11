package main

import (
	"context"
	"fmt"
	transfrom "github.com/aesoper101/kitextest/kitex_gen/transfrom"
)

// TransferServiceImpl implements the last service interface defined in the IDL.
type TransferServiceImpl struct{}

// TransIn implements the TransferServiceImpl interface.
func (s *TransferServiceImpl) TransIn(ctx context.Context, req *transfrom.TransferRequest) (resp *transfrom.Message, err error) {
	// TODO: Your code here...
	fmt.Println("TransIn", req)
	return &transfrom.Message{
		Reason: "TransIn success",
	}, nil
}

// TransOut implements the TransferServiceImpl interface.
func (s *TransferServiceImpl) TransOut(ctx context.Context, req *transfrom.TransferRequest) (resp *transfrom.Message, err error) {
	// TODO: Your code here...
	return
}
