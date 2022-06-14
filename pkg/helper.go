package pkg

import (
	"context"
	"sync"

	"github.com/houseme/gocrypto"
	"github.com/houseme/gocrypto/rsa"
)

var (
	once      sync.Once
	rsaSecret rsa.RSASecret
	insHelper = ihelper{}
)

// Helper Get Helper Instance
func Helper() *ihelper {
	return &insHelper
}

type ihelper struct {
}

// NewRsaSecretInstance Get RSASecret Instance
func NewRsaSecretInstance(ctx context.Context, publicKey, privateKey string, dataType gocrypto.Encode, keyType gocrypto.Secret) rsa.RSASecret {
	once.Do(func() {
		rsaSecret = rsa.RSASecret{
			PublicKey:          publicKey,
			PrivateKey:         privateKey,
			PublicKeyDataType:  dataType,
			PrivateKeyType:     keyType,
			PrivateKeyDataType: dataType,
		}
	})
	return rsaSecret
}

func (i *ihelper) HTTPPost(ctx context.Context, data interface{}) {

}
