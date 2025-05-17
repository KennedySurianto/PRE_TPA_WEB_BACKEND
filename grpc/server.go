package grpc

import (
	"fmt"
	"net"
	"pre_tpa_web/service"
	"pre_tpa_web/controller"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	grpcPort    int
	userService service.UserService
	server       *grpc.Server
}

func NewGrpcServer(grpcPort int, userService service.UserService) *GrpcServer {
	return &GrpcServer{
		grpcPort:     grpcPort,
		userService:  userService,
		server:       grpc.NewServer(),
	}
}

func (g *GrpcServer) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", g.grpcPort))
	if err != nil {
		panic(err)
	}
	controller.RegisterUserControllerServer(g.server, g.userService)
	err = g.server.Serve(listener)
	if err != nil {
		return
	}
}

func (g *GrpcServer) Stop() {
	g.server.GracefulStop()
}