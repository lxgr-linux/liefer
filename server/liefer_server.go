package server

import (
	"fmt"
	"github.com/lxgr-linux/liefer/build/project"
	"github.com/lxgr-linux/liefer/server/services"
	"github.com/lxgr-linux/liefer/server/types"
	"log"
)

type lieferServer struct {
	services.UnimplementedLieferServer
	pr *project.Registry
}

func (l *lieferServer) Deliver(payload *types.Payload, stream services.Liefer_DeliverServer) error {
	log.Println("hit")

	if payload.Body == nil {
		return fmt.Errorf("nil body received")
	}

	p, err := l.pr.GetProject(payload.Body.ProjectId)
	if err != nil {
		return err
	}

	err = p.Build(&stream)
	if err != nil {
		log.Printf("[Build][%s] failed: %s\n", payload.Body.ProjectId, err)
		return err
	}

	return nil
}
