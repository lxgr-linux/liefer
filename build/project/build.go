package project

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/lxgr-linux/liefer/server/services"
	"github.com/lxgr-linux/liefer/server/types"
)

func (p *Project) Build(branch string, stream *services.Liefer_DeliverServer) error {
	p.logger.Println("Building...")
	for {
		if !p.locked {
			break
		}
		err := (*stream).Send(types.ProgresNow(types.ProgressType_info, "Waiting for project to be freed...\n"))
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	p.locked = true
	defer p.unLock()

	err := p.setBranch(branch)
	if err != nil {
		return err
	}

	scriptFile, err := os.CreateTemp("", fmt.Sprintf("liefer-build-%s-*.sh", p.ID))
	if err != nil {
		return err
	}
	defer scriptFile.Close()
	defer os.Remove(scriptFile.Name())

	_, err = scriptFile.Write([]byte("set -o errexit\nset -o pipefail\nset -o verbose\n\n" + p.Script))
	if err != nil {
		return err
	}

	command := exec.Command("bash", scriptFile.Name())
	command.Dir = p.Location

	err = streamCommand(command, stream)
	if err != nil {
		p.logger.Printf("Failed to build: %s\n", err)
		return err
	}

	return nil
}

func streamCommand(command *exec.Cmd, stream *services.Liefer_DeliverServer) error {
	err := (*stream).Send(types.ProgresNow(types.ProgressType_info, "Starting build...\n"))
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

	var errCh = make(chan error)
	var wg sync.WaitGroup

	go readSteps(stdout, stream, types.ProgressType_info, errCh, &wg)
	go readSteps(stderr, stream, types.ProgressType_error, errCh, &wg)

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	}

	wg.Wait()
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

func readSteps(
	reader io.Reader, stream *services.Liefer_DeliverServer,
	progressType types.ProgressType,
	errCh chan error, wg *sync.WaitGroup,
) {
	var sendBuf []byte
	wg.Add(1)
	for {
		var buf = make([]byte, 128)
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			errCh <- err
			return
		}

		sendBuf = append(sendBuf, buf[:n]...)

		if strings.HasSuffix(string(sendBuf), "\n") {
			err = (*stream).Send(types.ProgresNow(progressType, string(sendBuf)))
			if err != nil {
				errCh <- err
				return
			}

			sendBuf = make([]byte, 0)
		}
	}
	wg.Done()
	errCh <- nil
}
