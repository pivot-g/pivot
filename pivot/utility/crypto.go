package utility

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/pivot-g/pivot/pivot/log"
)

type Key struct {
	Size           int
	Location       string
	PrivateKeyName string
	PublicKeyName  string
}

func (k *Key) GenPrivatePublicKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	privatekey, err := rsa.GenerateKey(rand.Reader, k.Size)
	if err != nil {
		log.Fatal("Cannot generate RSA key")
	}
	return privatekey, &privatekey.PublicKey
}

func (k *Key) WriteKey() {
	PrivateKey, PublicKey := k.GenPrivatePublicKey()

	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(PrivateKey),
	}
	privatePem, err := os.Create(fmt.Sprintf("%s/%s", k.Location, k.PrivateKeyName))
	if err != nil {
		log.Fatal("error when create private.pem: %s ", err)
	}
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		log.Fatal("error when create private.pem: %s ", err)
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(PublicKey)
	if err != nil {
		log.Fatal("error when dumping publickey: %s", err)
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create(fmt.Sprintf("%s/%s", k.Location, k.PublicKeyName))
	if err != nil {
		log.Fatal("error when create public.pem: %s ", err)
	}
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		log.Fatal("error when encode public pem: %s", err)
	}

}
