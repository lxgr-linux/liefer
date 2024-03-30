package project

import (
	"github.com/lxgr-linux/liefer/config"
	"github.com/lxgr-linux/liefer/server/services"
	"github.com/lxgr-linux/liefer/server/types"
	"os/exec"
)

type Project struct {
	config.Project
	locked bool
}

func (p *Project) unLock() {
	p.locked = false
}

func (p *Project) Build(stream *services.Liefer_DeliverServer) error {
	p.locked = true
	defer p.unLock()

	command := exec.Command(p.Script)

	err := (*stream).Send(types.ProgresNow("yes"))
	if err != nil {
		return err
	}

	err = command.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewProject(project config.Project) *Project {
	return &Project{project, false}
}
