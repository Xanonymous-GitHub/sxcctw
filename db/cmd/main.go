package main

import (
	"github.com/Xanonymous-GitHub/sxcctw/db/internal"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/proto"
	"github.com/Xanonymous-GitHub/sxcctw/db/pkg/vp"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func main() {
	serverAddress := ":" + strconv.Itoa(vp.Cvp.GetInt("serverAddress"))

	server := grpc.NewServer()

	proto.RegisterRecordServiceServer(server, &internal.RecordService{})

	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(lis)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}

	server.GracefulStop()
}
