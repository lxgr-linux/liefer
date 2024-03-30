package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"github.com/lxgr-linux/liefer/server/types"
	"google.golang.org/protobuf/proto"
)

func SignBody(body *types.Body, privKey *rsa.PrivateKey) ([]byte, error) {
	b, err := proto.Marshal(body)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(b)

	return rsa.SignPKCS1v15(nil, privKey, crypto.SHA256, hash[:])
}

func VerifyBody(body *types.Body, sig []byte, pubKey *rsa.PublicKey) error {
	b, err := proto.Marshal(body)
	if err != nil {
		return err
	}

	hash := sha256.Sum256(b)

	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], sig)
}

func NewPrivKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 512)
}

func PrivKeyToBase64(privKey *rsa.PrivateKey) []byte {
	privKeyBytes := x509.MarshalPKCS1PrivateKey(privKey)
	base64PrivKey := make([]byte, base64.StdEncoding.EncodedLen(len(privKeyBytes)))
	base64.StdEncoding.Encode(base64PrivKey, privKeyBytes)
	return base64PrivKey
}

func PrivKeyFromBase64(base64PrivKey []byte) (*rsa.PrivateKey, error) {
	privKeyBytes := make([]byte, base64.StdEncoding.DecodedLen(len(base64PrivKey)))
	_, err := base64.StdEncoding.Decode(privKeyBytes, base64PrivKey)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(privKeyBytes)
}
