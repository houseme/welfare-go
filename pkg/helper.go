package pkg

import (
	"context"
	"strings"
	"sync"

	"github.com/houseme/gocrypto"
	"github.com/houseme/gocrypto/rsa"
)

var (
	once      sync.Once
	rsaSecret rsa.RSASecret
	insHelper = ihelper{}
	// DefaultTrimChars are the characters which are stripped by Trim* functions in default.
	DefaultTrimChars = string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})
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

func (i *ihelper) Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.Trim(str, trimChars)
}
