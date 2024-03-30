package project

import (
	"fmt"
	"github.com/lxgr-linux/liefer/server/services"
	"github.com/lxgr-linux/liefer/server/types"
	"io"
	"os"
	"os/exec"
	"strings"
)

func (p *Project) Build(branch string, stream *services.Liefer_DeliverServer) error {
	p.logger.Println("Building...")
	p.locked = true
	defer p.unLock()

	err := p.setBranch(branch)
	if err != nil {
		return err
	}

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
	command.Dir = p.Location

	return streamCommand(command, stream)
}

func streamCommand(command *exec.Cmd, stream *services.Liefer_DeliverServer) error {
	err := (*stream).Send(types.ProgresNow(types.ProgressType_info, "Starting build...\n"))
	if err != nil {
		return err
	}

	output, err := getCommandOutput(command)
	err = command.Start()
	if err != nil {
		return err
	}

	var sendBuf []byte
	for {
		var buf = make([]byte, 128)
		n, err := output.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		sendBuf = append(sendBuf, buf[:n]...)

		if strings.HasSuffix(string(sendBuf), "\n") {
			err = (*stream).Send(types.ProgresNow(types.ProgressType_info, string(sendBuf)))
			if err != nil {
				return err
			}

			sendBuf = make([]byte, 0)
		}
	}

	err = command.Wait()
	if err != nil {
		err = (*stream).Send(
			types.ProgresNow(
				types.ProgressType_error,
				err.Error(),
			),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func getCommandOutput(command *exec.Cmd) (io.Reader, error) {
	stdout, err := command.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		return nil, err
	}

	return io.MultiReader(stdout, stderr), nil
}
