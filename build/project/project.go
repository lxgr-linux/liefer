package project

import (
	"fmt"
	"github.com/lxgr-linux/liefer/config"
	"github.com/lxgr-linux/liefer/server/services"
	"github.com/lxgr-linux/liefer/server/types"
	"io"
	"log"
	"os"
	"os/exec"
)

type Project struct {
	config.Project
	logger *log.Logger
	locked bool
}

func (p *Project) unLock() {
	p.locked = false
}

func (p *Project) Build(stream *services.Liefer_DeliverServer) error {
	p.locked = true
	defer p.unLock()

	scriptFile, err := os.CreateTemp("/tmp", fmt.Sprintf("liefer-build-%s-*.sh", p.ID))
	if err != nil {
		return err
	}
	defer scriptFile.Close()
	defer os.Remove(scriptFile.Name())

	_, err = scriptFile.Write([]byte(p.Script))
	if err != nil {
		return err
	}

	command := exec.Command("sh", scriptFile.Name())

	err = (*stream).Send(types.ProgresNow(types.ProgressType_info, "Starting build...\n"))
	if err != nil {
		return err
	}

	stdout, err := command.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		return err
	}

	err = command.Start()
	if err != nil {
		return err
	}

	for {
		var buf = make([]byte, 32)
		n, err := stdout.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		err = (*stream).Send(types.ProgresNow(types.ProgressType_info, string(buf[:n])))
		if err != nil {
			return err
		}

	}

	errorData, err := io.ReadAll(stderr)
	if err != nil {
		return err
	}

	err = command.Wait()
	if err != nil {
		err = (*stream).Send(
			types.ProgresNow(
				types.ProgressType_error,
				fmt.Sprintf("%s: %s", err, string(errorData)),
			),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewProject(project config.Project) *Project {
	return &Project{
		project,
		log.New(
			os.Stdout,
			fmt.Sprintf("[%s][Build]: ", project.ID),
			log.Lmsgprefix|log.LstdFlags),
		false}
}
