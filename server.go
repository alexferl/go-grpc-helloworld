package go_grpc_helloworld

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/alexferl/go_grpc_helloworld/methods"
)

var (
	grpcServer       *grpc.Server
	grpcHealthServer *grpc.Server
	system           = "" // empty string represents the health of the system
)

func init() {
	c := NewConfig()
	c.BindFlags()
}

func Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sig)

	g, ctx := errgroup.WithContext(ctx)

	// gRPC Health Server
	healthServer := health.NewServer()
	g.Go(func() error {
		addr := fmt.Sprintf("%s:%s", viper.GetString("health-bind-address"), viper.GetString("health-bind-port"))
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal().Msgf("health server failed to listen: '%v'", err)
		}

		grpcHealthServer = grpc.NewServer()
		healthpb.RegisterHealthServer(grpcHealthServer, healthServer)
		if viper.GetString("env-name") == "local" {
			reflection.Register(grpcHealthServer)
		}

		log.Info().Msgf("gRPC health server serving at: %s", addr)
		return grpcHealthServer.Serve(lis)
	})

	// gRPC server
	g.Go(func() error {
		addr := fmt.Sprintf("%s:%d", viper.GetString("bind-address"), viper.GetInt("bind-port"))
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal().Msgf("server failed to listen: '%v'", err)
		}

		grpcServer = grpc.NewServer(
			grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: 2 * time.Minute}),
		)
		pb.RegisterGreeterServer(grpcServer, &methods.Server{})
		if viper.GetString("env-name") == "local" {
			reflection.Register(grpcServer)
		}

		log.Info().Msgf("gRPC server serving at %s", addr)
		healthServer.SetServingStatus(system, healthpb.HealthCheckResponse_SERVING)
		return grpcServer.Serve(lis)
	})

	select {
	case <-sig:
		break
	case <-ctx.Done():
		break
	}

	cancel()

	healthServer.SetServingStatus(system, healthpb.HealthCheckResponse_NOT_SERVING)

	timeout := time.Duration(viper.GetInt64("graceful-timeout")) * time.Second
	_, shutdownCancel := context.WithTimeout(context.Background(), timeout)
	defer shutdownCancel()

	if grpcServer != nil {
		grpcServer.GracefulStop()
	}
	if grpcHealthServer != nil {
		grpcHealthServer.GracefulStop()
	}

	err := g.Wait()
	if err != nil {
		log.Fatal().Msgf("wait returned error: '%v'", err)
	}
}
