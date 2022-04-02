package main

import (
	"github.com/Xanonymous-GitHub/sxcctw/pkg/ginx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			ginx.NewEngine,
		),
		fx.Invoke(),
	).Run()
}
