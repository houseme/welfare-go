package welfare

import (
	"context"

	"github.com/houseme/gocrypto"
	"github.com/houseme/gocrypto/rsa"

	"github.com/houseme/welfare-go/internal"
	"github.com/houseme/welfare-go/pkg"
	"github.com/houseme/welfare-go/pkg/log"
)

// Welfare .
type Welfare struct {
	welfare *internal.Welfare
}

// New .create a new Welfare.
func New(ctx context.Context, channel, environment string, opts ...internal.Option) (*Welfare, error) {
	if len(pkg.Helper().Trim(channel)) == 0 {
		return nil, internal.ErrInvalidChannel
	}

	if len(pkg.Helper().Trim(environment)) == 0 {
		return nil, internal.ErrInvalidEnvironment
	}

	return &Welfare{
		welfare: internal.New(ctx, channel, environment, opts...),
	}, nil
}

// Init .
func (w *Welfare) Init(ctx context.Context, opts ...internal.Option) error {
	if len(opts) < 2 {
		return internal.ErrInvalidOptionLength
	}
	return w.welfare.Init(ctx, opts...)
}

// InitSecret .
func (w *Welfare) InitSecret(ctx context.Context, publicKey, privateKey string) error {
	if len(publicKey) == 0 || len(privateKey) == 0 {
		return internal.ErrInvalidPublicKeyOrPrivateKey
	}
	w.welfare.SetPublicKeyAndPrivateKey(ctx, publicKey, privateKey)
	return w.welfare.InitSecret(ctx)
}

// SetPublicKey .
func (w *Welfare) SetPublicKey(ctx context.Context, publicKey string) {
	w.welfare.SetPublicKey(ctx, publicKey)
}

// SetPrivateKey .
func (w *Welfare) SetPrivateKey(ctx context.Context, privateKey string) {
	w.welfare.SetPrivateKey(ctx, privateKey)
}

// SetPublicKeyDataType .
func (w *Welfare) SetPublicKeyDataType(ctx context.Context, dataType gocrypto.Encode) {
	w.welfare.SetPublicKeyDataType(ctx, dataType)
}

//	SetPrivateKeyDataType .
func (w *Welfare) SetPrivateKeyDataType(ctx context.Context, dataType gocrypto.Encode) {
	w.welfare.SetPrivateKeyDataType(ctx, dataType)
}

//	SetPrivateKeyType .
func (w *Welfare) SetPrivateKeyType(ctx context.Context, keyType gocrypto.Secret) {
	w.welfare.SetPrivateKeyType(ctx, keyType)
}

// SetPublicKeyAndPrivateKey .
func (w *Welfare) SetPublicKeyAndPrivateKey(ctx context.Context, publicKey, privateKey string) {
	w.welfare.SetPublicKeyAndPrivateKey(ctx, publicKey, privateKey)
}

// SetLogger .
func (w *Welfare) SetLogger(ctx context.Context, logger *log.Logger) {
	w.welfare.SetLogger(ctx, logger)
}

// SetChannel .
func (w *Welfare) SetChannel(ctx context.Context, channel string) {
	w.welfare.SetChannel(ctx, channel)
}

// GetSecret get secret.
func (w *Welfare) GetSecret(ctx context.Context) (secret rsa.RSASecret, err error) {
	return w.welfare.Secret(ctx)
}
