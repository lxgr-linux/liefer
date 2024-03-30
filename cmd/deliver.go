package cmd

import (
	"fmt"
	"github.com/lxgr-linux/liefer/client"
	"github.com/lxgr-linux/liefer/crypto"
	"github.com/lxgr-linux/liefer/server/types"
	"github.com/spf13/cobra"
)

var address string

func init() {
	deliverCmd.Flags().StringVarP(&address, "address", "a", "localhost:8080", "the address of the liefer server")
	rootCmd.AddCommand(deliverCmd)
}

var deliverCmd = &cobra.Command{
	Use:   "deliver [project-id] [branch] [priv-key]",
	Short: "delivers to the remote liefer instance",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectId := args[0]
		if projectId == "" {
			return fmt.Errorf("projectId is empty")
		}

		branch := args[1]
		if branch == "" {
			return fmt.Errorf("branch is empty")
		}

		privKey, err := crypto.PrivKeyFromBase64([]byte(args[2]))
		if err != nil {
			return err
		}

		body := types.Body{ProjectId: projectId, Branch: branch}

		sig, err := crypto.SignBody(&body, privKey)
		if err != nil {
			return err
		}

		client, err := client.Connect(address)
		if err != nil {
			return err
		}
		defer client.Disconnect()

		return client.SendDeliver(&types.Payload{Body: &body, Sig: sig})
	},
}
