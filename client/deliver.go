package client

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/lxgr-linux/liefer/server/types"
)

func (c Client) SendDeliver(payload *types.Payload) error {
	var escape string
	stream, err := c.client.Deliver(context.Background(), payload)
	if err != nil {
		return err
	}

	for {
		progress, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if progress.Type == types.ProgressType_error {
			escape = "\x1b[1;31m"
		} else {
			escape = ""
		}

		timeStamp := time.Unix(progress.Timestamp, 0).UTC().Format(time.DateTime)

		for _, line := range strings.Split(progress.Content, "\n") {
			if line != "" {
				fmt.Printf(
					"%s\t%s%s\x1b[0m\n",
					timeStamp,
					escape,
					line,
				)
			}
		}
	}
	if escape != "" {
		os.Exit(1)
	}

	return nil
}
