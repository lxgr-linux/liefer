package project

import (
	"fmt"
	"github.com/lxgr-linux/liefer/config"
)

type Registry struct {
	projects []*Project
}

func (pr *Registry) Register(p *Project) error {
	_, err := pr.GetProject(p.ID)
	if err == nil {
		return fmt.Errorf("project `%s` already registred", p.ID)
	}
	pr.projects = append(pr.projects, p)

	return nil
}

func (pr *Registry) GetProject(id string) (*Project, error) {
	for _, p := range pr.projects {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, fmt.Errorf("project `%s` not found", id)
}

func RegistryFromConfig(cfg *config.Config) (*Registry, error) {
	var pr Registry

	for _, project := range cfg.Projects {
		err := pr.Register(NewProject(project))
		if err != nil {
			return nil, err
		}
	}

	return &pr, nil
}
