package project

import (
    "fmt"
    "github.com/lxgr-linux/liefer/build/git"
    "github.com/lxgr-linux/liefer/config"
    "github.com/lxgr-linux/liefer/logger"
    "log"
    "os"
)

type Project struct {
    *config.Project
    logger *log.Logger
    locked bool
    git    *git.Git
}

func (p *Project) CreateDir() error {
    stat, err := os.Stat(p.Location)

    if err != nil {
        err = os.MkdirAll(p.Location, os.ModePerm)
        if err != nil {
            return err
        }

        err = p.git.Clone()
        if err != nil {
            return err
        }
    } else if !stat.IsDir() {
        return fmt.Errorf("%s exists, but is not a dir", p.Location)
    } else {
        err = p.git.Pull()
        if err != nil {
            return err
        }
    }

    return nil
}

func (p *Project) unLock() {
    p.locked = false
}

func (p *Project) setBranch(branch string) error {
    err := p.git.ResetDir()
    if err != nil {
        return err
    }

    err = p.git.Fetch()
    if err != nil {
        return err
    }

    err = p.git.Checkout(branch)
    if err != nil {
        return err
    }

    return p.git.Pull()
}

func NewProject(project *config.Project) *Project {
    return &Project{
        project,
        logger.Build(fmt.Sprintf("[%s][Build]", project.ID)),
        false,
        git.NewGit(project.ID, project.Location, project.Repo),
    }
}
