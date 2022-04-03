package main

import (
	"github.com/Xanonymous-GitHub/sxcctw/internal/api/client"
	"github.com/Xanonymous-GitHub/sxcctw/internal/api/server"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/ginx"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/logrux"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			ginx.NewEngine,
			logrux.NewLogger,
		),
		fx.Invoke(
			server.RegisterRestApiHandlers,
			client.NewRecordServiceClient,
		),
	).Run()
}
