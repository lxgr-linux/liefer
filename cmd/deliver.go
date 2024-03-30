package cmd

import (
	"fmt"
	"github.com/lxgr-linux/liefer/client"
	"github.com/lxgr-linux/liefer/server/types"
	"github.com/spf13/cobra"
)

var address string

func init() {
	deliverCmd.Flags().StringVarP(&address, "address", "a", "localhost:8080", "the address of the liefer server")
	rootCmd.AddCommand(deliverCmd)
}

var deliverCmd = &cobra.Command{
	Use:   "deliver [project-id] [branch]",
	Short: "delivers to the remote liefer instance",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectId := args[0]
		if projectId == "" {
			return fmt.Errorf("projectId is empty")
		}

		branch := args[1]
		if branch == "" {
			return fmt.Errorf("branch is empty")
		}

		body := types.Body{ProjectId: projectId, Branch: branch}

		client, err := client.Connect(address)
		if err != nil {
			return err
		}
		defer client.Disconnect()

		return client.SendDeliver(&types.Payload{Body: &body})
	},
}
