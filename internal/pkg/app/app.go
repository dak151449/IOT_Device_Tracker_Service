package app

import (
	"context"
	"fmt"
	"iot-device-tracker-service/internal/pkg/app/mw"
	"iot-device-tracker-service/internal/pkg/config"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcListener net.Listener
	closer       chan struct{}
	grpcServer   *grpc.Server
}

func New() (*App, error) {
	a := &App{}

	lsn, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GetStringValue(config.GrpcPORT)))
	if err != nil {
		return nil, err
	}

	a.grpcListener = lsn

	a.InitCloser(syscall.SIGTERM, syscall.SIGINT)

	return a, nil
}

type Service interface {
	RegisterGRPC(*grpc.Server)
}

func (a *App) Run(options []grpc.UnaryServerInterceptor, impl ...Service) error {
	opts := []grpc.UnaryServerInterceptor{mw.LogResponseInterceptor}
	opts = append(opts, options...)
	a.grpcServer = grpc.NewServer(grpc.ChainUnaryInterceptor(opts...))

	for _, i := range impl {
		i.RegisterGRPC(a.grpcServer)
	}
	reflection.Register(a.grpcServer)

	go func() {
		if err := a.grpcServer.Serve(a.grpcListener); err != nil {
			log.Error().Err(err).Msg("gRPC server, error when starting")
			a.ShutDown()
		}
	}()

	strBuilder := &strings.Builder{}
	strBuilder.WriteString("App started. Ports: ")
	grpcPort := a.grpcListener.Addr().(*net.TCPAddr).Port
	strBuilder.WriteString(fmt.Sprintf("GRPC - %d", grpcPort))
	log.Info().Msg(strBuilder.String())

	a.Wait()

	return nil
}

func (a *App) InitCloser(sig ...os.Signal) {
	a.closer = make(chan struct{})
	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			a.ShutDown()
		}()
	}
}

func (a *App) Wait() {
	<-a.closer
}

func (a *App) ShutDown() {
	defer close(a.closer)
	log.Info().Msg("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	done := make(chan struct{})
	go func() {
		a.grpcServer.GracefulStop()
		close(done)
	}()
	select {
	case <-done:
		log.Info().Msg("gRPC server gracefully stopped")
	case <-ctx.Done():
		err := fmt.Errorf("error when shutting down: %w", ctx.Err())
		a.grpcServer.Stop()
		log.Error().Err(err).Msg("gRPC server force stopped")
	}
}
