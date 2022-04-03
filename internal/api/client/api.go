package client

import (
	"context"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/env"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/grpc"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/proto/pb"
	"go.uber.org/fx"
)

func NewRecordServiceClient(lc fx.Lifecycle) (pb.RecordServiceClient, error) {
	conn, err := grpc.NewGRPCConn(env.DBGrpcServerHost + ":" + env.DBGrpcServerPort)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			err := conn.Close()
			if err != nil {
				return err
			}

			return nil
		},
	})

	return pb.NewRecordServiceClient(conn), nil
}
