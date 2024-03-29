package cmd

import (
	"github.com/lxgr-linux/liefer/config"
	"github.com/lxgr-linux/liefer/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve [config.yaml]",
	Short: "starts service",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		cfgPath := "./config.yaml"
		if len(args) == 1 {
			cfgPath = args[0]
		}

		cfg, err := config.Read(cfgPath)
		if err != nil {
			return err
		}

		return server.Serve(cfg)
	},
}
