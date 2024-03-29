package server

import (
	"github.com/lxgr-linux/liefer/server/services"
	"github.com/lxgr-linux/liefer/server/types"
	"log"
)

type lieferServer struct {
	services.UnimplementedLieferServer
}

func (l *lieferServer) Deliver(payload *types.Payload, stream services.Liefer_DeliverServer) error {
	log.Println("hit")

	return nil
}
