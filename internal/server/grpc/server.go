package grpcserver

import (
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
	// TODO: add logger and app
}

func NewServer(app *app.App) *RotatorServer {
	return &RotatorServer{
		App: app,
	}
}

func (r *RotatorServer) Start(network string, address string) error {
	listen, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	r.Server = grpc.NewServer()
	rotatorpb.RegisterRotatorServer(r.Server, r)
	fmt.Println("server is running")
	return r.Server.Serve(listen)
}

func (r *RotatorServer) Stop() {
	r.Server.GracefulStop()
}
