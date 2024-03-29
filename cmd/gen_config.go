package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/lxgr-linux/liefer/config"
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

		privKey, err := rsa.GenerateKey(rand.Reader, 512)
		if err != nil {
			return err
		}
		privKeyBytes := x509.MarshalPKCS1PrivateKey(privKey)
		base64PrivKey := make([]byte, base64.StdEncoding.EncodedLen(len(privKeyBytes)))
		base64.StdEncoding.Encode(base64PrivKey, privKeyBytes)

		fmt.Printf("Private key:\n%s\n", string(base64PrivKey))

		cfg := config.Config{PubKey: x509.MarshalPKCS1PublicKey(&privKey.PublicKey), Host: "localhost", Port: 8080}

		return cfg.Write(cfgPath)
	},
}
