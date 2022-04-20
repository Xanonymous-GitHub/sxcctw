package ginx

import (
	"context"
	"fmt"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/fx"
	"log"
	"net/http"
)

// NewEngine return *gin.Engine for gin router
func NewEngine(lc fx.Lifecycle) (*gin.Engine, error) {
	if env.IsDebugMode {
		gin.ForceConsoleColor()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	engine.Use(
		otelgin.Middleware("gin"),
		gin.Logger(),
		gin.Recovery(),
	)

	if env.IsDebugMode {
		engine.Use(cors.Default())
	}

	engine.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	engine.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	serverAddr := fmt.Sprintf(":%s", env.ApiRestServerPort)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					panic(err)
				}
			}()
			log.Printf("[GIN] server start and listen (%s)", serverAddr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Print("[GIN] server graceful stop")
			return srv.Shutdown(ctx)
		},
	})
	return engine, nil
}
