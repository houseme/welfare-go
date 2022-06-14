package welfare

import (
	"context"
	"time"

	"github.com/houseme/gocrypto"

	"github.com/houseme/welfare-go/internal"
	"github.com/houseme/welfare-go/pkg"
)

// Welfare .
type Welfare struct {
	welfare *internal.Welfare
}

// New .create
func New(ctx context.Context, environment string, opts ...internal.Option) *Welfare {
	var gateway = internal.DevGateway
	if environment == internal.ProdEnvironment {
		gateway = internal.ProdGateway
	}

	op := internal.Options{
		Version:            internal.Version,
		Channel:            "",
		Format:             internal.Format,
		EncryptType:        internal.EncryptType,
		SignType:           internal.SignType,
		PublicKey:          "",
		PublicKeyDataType:  internal.PublicKeyDataType,
		PrivateKeyDataType: internal.PrivateKeyDataType,
		PrivateKey:         "",
		PrivateKeyType:     internal.PrivateKeyType,
		Timeout:            internal.Timeout,
		RetryCount:         0,
	}
	for _, option := range opts {
		option(&op)
	}

	return &Welfare{welfare: &internal.Welfare{
		Gateway: gateway,
		Logger:  pkg.NewLog(ctx, pkg.LogPath, pkg.DebugLevel),
		Opt:     op,
	}}
}

// WithVersion .
func WithVersion(version string) internal.Option {
	return func(o *internal.Options) {
		o.Version = version
	}
}

// WithChannel .
func WithChannel(channel string) internal.Option {
	return func(o *internal.Options) {
		o.Channel = channel
	}
}

// WithFormat .
func WithFormat(format string) internal.Option {
	return func(o *internal.Options) {
		o.Format = format
	}
}

// WithEncryptType .
func WithEncryptType(encryptType string) internal.Option {
	return func(o *internal.Options) {
		o.EncryptType = encryptType
	}
}

// WithSignType .
func WithSignType(signType string) internal.Option {
	return func(o *internal.Options) {
		o.SignType = signType
	}
}

// WithPublicKey .
func WithPublicKey(publicKey string) internal.Option {
	return func(o *internal.Options) {
		o.PublicKey = publicKey
	}
}

// WithPrivateKey .
func WithPrivateKey(privateKey string) internal.Option {
	return func(o *internal.Options) {
		o.PrivateKey = privateKey
	}
}

// WithPublicKeyDataType .
func WithPublicKeyDataType(dataType gocrypto.Encode) internal.Option {
	return func(o *internal.Options) {
		o.PublicKeyDataType = dataType
	}
}

// WithPrivateKeyDataType .
func WithPrivateKeyDataType(dataType gocrypto.Encode) internal.Option {
	return func(o *internal.Options) {
		o.PrivateKeyDataType = dataType
	}
}

// WithPrivateKeyType .
func WithPrivateKeyType(keyType gocrypto.Secret) internal.Option {
	return func(o *internal.Options) {
		o.PrivateKeyType = keyType
	}
}

// WithTimeout the Timeout option.
func WithTimeout(timeout time.Duration) internal.Option {
	return func(o *internal.Options) {
		o.Timeout = timeout
	}
}

// WithRetryCount with RetryCount option.
func WithRetryCount(retryCount int) internal.Option {
	return func(o *internal.Options) {
		o.RetryCount = retryCount
	}
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

// InitRSA .
func (w *Welfare) InitRSA(ctx context.Context, publicKey, privateKey string) error {
	if len(publicKey) == 0 || len(privateKey) == 0 {
		return pkg.ErrInvalidKey
	}
	return w.welfare.InitRSA(ctx)
}

// SetPublicKeyAndPrivateKey .
func (w *Welfare) SetPublicKeyAndPrivateKey(ctx context.Context, publicKey, privateKey string) {
	w.welfare.SetPublicKeyAndPrivateKey(ctx, publicKey, privateKey)
}

// SetLogger .
func (w *Welfare) SetLogger(ctx context.Context, logger *pkg.Logger) {
	w.welfare.SetLogger(ctx, logger)
}

// SetChannel .
func (w *Welfare) SetChannel(ctx context.Context, channel string) {
	w.welfare.SetChannel(ctx, channel)
}
