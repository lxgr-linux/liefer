package client

import (
	"context"
	"fmt"
	"github.com/lxgr-linux/liefer/server/types"
	"io"
)

func (c Client) SendDeliver(payload *types.Payload) error {
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

		fmt.Printf("%v\n", progress)
	}

	return nil
}
