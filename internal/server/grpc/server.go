package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/fevse/banners_rotator/internal/app"
	rotatorpb "github.com/fevse/banners_rotator/internal/server/grpc/pb"
	"google.golang.org/grpc"
)

type RotatorServer struct {
	rotatorpb.UnimplementedRotatorServer
	Server *grpc.Server
	App    *app.App
}

type Logger interface {
	Info(string)
	Error(string)
}

func NewServer(app *app.App) *RotatorServer {
	return &RotatorServer{
		App: app,
	}
}

func (r *RotatorServer) Start(network string, address string) error {
	listen, err := net.Listen(network, address)
	if err != nil {
		r.App.Logger.Error("failed to start grpcserver" + err.Error())
	}
	r.Server = grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor(r.App.Logger)))
	rotatorpb.RegisterRotatorServer(r.Server, r)
	r.App.Logger.Info("server is running: " + address)
	return r.Server.Serve(listen)
}

func (r *RotatorServer) Stop() {
	r.App.Logger.Info("server stopped")
	r.Server.GracefulStop()
}

func loggingInterceptor(logg app.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		logg.Info(info.FullMethod + " " + fmt.Sprintf("%v", req))
		return handler(ctx, req)
	}
}
