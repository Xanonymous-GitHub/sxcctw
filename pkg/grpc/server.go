package grpc

import (
	"context"
	"fmt"
	"github.com/Xanonymous-GitHub/sxcctw/pkg/env"
	gram "github.com/grpc-ecosystem/go-grpc-middleware"
	glorious "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	greco "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	promethium "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	otl "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthcare "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"time"
)

func NewGRPCServer(lc fx.Lifecycle) *grpc.Server {
	grpcRecoveryOpts := []greco.Option{
		greco.WithRecoveryHandler(func(p interface{}) (err error) {
			fmt.Printf("panic: %v\n", p)
			debug.PrintStack()
			return status.Errorf(codes.Internal, "panic: %v", p)
		}),
	}

	promethium.EnableHandlingTimeHistogram()

	server := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             5 * time.Second,
				PermitWithoutStream: true,
			},
		),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle:     15 * time.Second,
				MaxConnectionAge:      30 * time.Second,
				MaxConnectionAgeGrace: 5 * time.Second,
				Time:                  5 * time.Second,
				Timeout:               1 * time.Second,
			},
		),
		grpc.UnaryInterceptor(gram.ChainUnaryServer(
			otl.UnaryServerInterceptor(),
			promethium.UnaryServerInterceptor,
			glorious.UnaryServerInterceptor(logrus.NewEntry(logrus.New())),
			greco.UnaryServerInterceptor(grpcRecoveryOpts...),
		)),
		grpc.StreamInterceptor(gram.ChainStreamServer(
			otl.StreamServerInterceptor(),
			promethium.StreamServerInterceptor,
			glorious.StreamServerInterceptor(logrus.NewEntry(logrus.New())),
			greco.StreamServerInterceptor(grpcRecoveryOpts...),
		)),
	)

	healthServer := health.NewServer()
	healthcare.RegisterHealthServer(server, healthServer)
	reflection.Register(server)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			serverAddr := fmt.Sprintf(":%s", env.DBGrpcServerPort)
			go func() {
				lis, err := net.Listen("tcp", serverAddr)
				if err != nil {
					panic(err)
				}
				if err := server.Serve(lis); err != nil {
					panic(err)
				}
			}()
			log.Printf("[gRPC-Go] server start and listen (%s)", serverAddr)
			metricsAddr := fmt.Sprintf(":%s", env.DBGrpcServerMetricsPort)
			go func() {
				http.Handle("/metrics", promhttp.Handler())
				_ = http.ListenAndServe(metricsAddr, nil)
			}()
			log.Printf("[gRPC-Go] metrics server start and listen (%s)", metricsAddr)
			return nil
		},
		OnStop: func(context.Context) error {
			healthServer.SetServingStatus("", healthcare.HealthCheckResponse_NOT_SERVING)
			log.Print("[gRPC-Go] set health check response: NOT_SERVING")
			server.GracefulStop()
			log.Print("[gRPC-Go] server graceful stop")
			return nil
		},
	})

	return server
}
