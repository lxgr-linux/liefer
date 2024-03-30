package server

import (
	"fmt"
	"github.com/lxgr-linux/liefer/build/project"
	"github.com/lxgr-linux/liefer/config"
	"github.com/lxgr-linux/liefer/server/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Serve(pr *project.Registry, cfg *config.Config) error {
	host := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	services.RegisterLieferServer(grpcServer, &lieferServer{pr: pr})

	log.Printf("Starting grpc server at %s...\n", host)

	return grpcServer.Serve(lis)
}
