package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "liefer",
	Short: "Liefer cd platform",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("no arguments given")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
