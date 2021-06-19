package server

import (
	v1 "github.com/llwwbb/geektime_practice/go_advanced_training/week4/api/user/service/v1"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/conf"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/service"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	*grpc.Server
	c *conf.Server
}

func NewGrpcServer(c *conf.Server, s *service.UserService) *GrpcServer {
	server := grpc.NewServer()
	server.RegisterService(&v1.User_ServiceDesc, s)
	return &GrpcServer{c: c, Server: server}
}

func (s *GrpcServer) Start() error {
	lis, err := net.Listen("tcp", s.c.Http.Addr)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

func (s *GrpcServer) Stop() error {
	s.GracefulStop()
	return nil
}
