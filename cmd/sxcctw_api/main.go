package main

import (
	"github.com/Xanonymous-GitHub/sxcctw/internal/api/client"
	"github.com/Xanonymous-GitHub/sxcctw/internal/api/handler"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/ginx"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/logrux"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			ginx.NewEngine,
			logrux.NewLogger,
			client.NewRecordServiceClient,
		),
		fx.Invoke(
			handler.RegisterRestApiHandlers,
		),
	).Run()
}
