package client

import (
	"github.com/lxgr-linux/liefer/server/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client services.LieferClient
	conn   *grpc.ClientConn
}

func (c Client) Disconnect() {
	c.conn.Close()
}

func Connect(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{services.NewLieferClient(conn), conn}, nil
}
