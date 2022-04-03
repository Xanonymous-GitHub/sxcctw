package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCConn(serverAddr string) (*grpc.ClientConn, error) {
	var options []grpc.DialOption
	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.Dial(serverAddr, options...)
}
