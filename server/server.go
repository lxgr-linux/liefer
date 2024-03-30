package server

import (
	"crypto/x509"
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

	pubKey, err := x509.ParsePKCS1PublicKey(cfg.PubKey)
	if err != nil {
		return err
	}
	services.RegisterLieferServer(grpcServer, &lieferServer{pr: pr, pubKey: pubKey})

	log.Printf("Starting grpc server at %s...\n", host)

	return grpcServer.Serve(lis)
}
