package cmd

import (
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
	Use:   "deliver",
	Short: "delivers to the remote liefer instance",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		client, err := client.Connect(address)
		if err != nil {
			return err
		}
		defer client.Disconnect()

		return client.SendDeliver(&types.Payload{})
	},
}
