package main

import (
	"github.com/Xanonymous-GitHub/sxcctw/internal/api/server"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/gorm"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/grpc"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/logrux"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			grpc.NewGRPCServer,
			logrux.NewLogger,
			gorm.NewGORMMySQLClient,
			gorm.CreateNewDBClientWith,
		),
		fx.Invoke(
			server.RegisterRecordServiceServer,
			gorm.InitializeDBTables,
		),
	).Run()
}
