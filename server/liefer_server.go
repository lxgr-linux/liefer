package server

import (
	"fmt"
	"github.com/lxgr-linux/liefer/build/project"
	"github.com/lxgr-linux/liefer/server/services"
	"github.com/lxgr-linux/liefer/server/types"
)

type lieferServer struct {
	services.UnimplementedLieferServer
	pr *project.Registry
}

func (l *lieferServer) Deliver(payload *types.Payload, stream services.Liefer_DeliverServer) error {
	if payload.Body == nil {
		return fmt.Errorf("nil body received")
	}

	p, err := l.pr.GetProject(payload.Body.ProjectId)
	if err != nil {
		return err
	}

	err = p.Build(payload.Body.Branch, &stream)
	if err != nil {
		return err
	}

	return nil
}
