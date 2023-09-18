package app

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"iot-device-tracker-service/internal/config"
	"iot-device-tracker-service/internal/pkg/mw"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type App struct {
	grpcListener net.Listener
	closer       chan struct{}
	grpcServer   *grpc.Server
}

func New() (*App, error) {
	a := &App{}

	lsn, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GetValue(config.GrpcPORT)))
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

func (a *App) Run(impl ...Service) error {
	a.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(mw.LogResponseInterceptor))

	for _, i := range impl {
		i.RegisterGRPC(a.grpcServer)
	}
	reflection.Register(a.grpcServer)

	go func() {
		if err := a.grpcServer.Serve(a.grpcListener); err != nil {
			log.Error().Err(err).Msg("gRPC server")
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

func (a *App) ShutDown() {
	defer close(a.closer)
	log.Info().Msg("shutting down...")
}

func (a *App) Wait() {
	<-a.closer
}
