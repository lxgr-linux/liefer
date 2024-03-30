package cmd

import (
	"crypto/x509"
	"fmt"
	"github.com/lxgr-linux/liefer/config"
	"github.com/lxgr-linux/liefer/crypto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(genConfigCmd)
}

var genConfigCmd = &cobra.Command{
	Use:   "gen-config [config.yaml]",
	Short: "generates a config",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgPath := "./config.yaml"
		if len(args) == 1 {
			cfgPath = args[0]
		}

		privKey, err := crypto.NewPrivKey()
		if err != nil {
			return err
		}

		fmt.Printf("Private key:\n%s\n", crypto.PrivKeyToString(privKey))

		cfg := config.Config{PubKey: x509.MarshalPKCS1PublicKey(&privKey.PublicKey), Host: "localhost", Port: 8080}

		return cfg.Write(cfgPath)
	},
}
