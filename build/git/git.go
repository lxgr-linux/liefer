package git

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/lxgr-linux/liefer/logger"
)

type Git struct {
	logger *log.Logger
	dir    string
	repo   string
}

func (g Git) Clone() error {
	g.logger.Printf("Cloning %s to %s...\n", g.repo, g.dir)
	return exec.Command("git", "clone", g.repo, g.dir).Run()
}

func (g Git) Fetch() error {
	g.logger.Println("Fetching...")

	command := exec.Command("git", "fetch")
	command.Dir = g.dir

	return command.Run()
}

func (g Git) Pull() error {
	g.logger.Println("Pull...")

	command := exec.Command("git", "pull")
	command.Dir = g.dir

	return command.Run()
}

func (g Git) Checkout(branch string) error {
	g.logger.Printf("Checking out %s...\n", branch)

	command := exec.Command("git", "checkout", branch)
	command.Dir = g.dir

	return command.Run()
}

func (g Git) ResetDir() error {
	g.logger.Println("Resetting...")

	checkoutCmd := exec.Command("git", "checkout", ".")
	checkoutCmd.Dir = g.dir
	err := checkoutCmd.Run()
	if err != nil {
		return err
	}

	cleanCmd := exec.Command("git", "clean", "-df")
	cleanCmd.Dir = g.dir
	return cleanCmd.Run()
}

func NewGit(id, dir, repo string) *Git {
	return &Git{
		logger: logger.Build(fmt.Sprintf("[%s][git]", id)),
		dir:    dir,
		repo:   repo,
	}
}
