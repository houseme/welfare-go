package internal

import (
	"context"
	"os"

	"github.com/houseme/gocrypto/rsa"

	"github.com/houseme/welfare-go/pkg/log"
)

// Welfare .
type Welfare struct {
	ctx         context.Context
	gateway     string
	channel     string
	serviceType string
	opt         Options
	logger      *log.Logger
	secret      *rsa.RSASecret
}

// New Create a new Welfare.
func New(ctx context.Context, channel, environment string, opts ...Option) *Welfare {
	var gateway = devGateway
	if environment == prodEnvironment {
		gateway = prodGateway
	}

	op := Options{
		Version:            version,
		Format:             format,
		EncryptType:        encryptType,
		SignType:           signType,
		PublicKey:          "",
		PublicKeyDataType:  publicKeyDataType,
		PrivateKeyDataType: privateKeyDataType,
		PrivateKey:         "",
		PrivateKeyType:     privateKeyType,
		Timeout:            timeout,
		RetryCount:         0,
		Loglevel:           log.InfoLevel,
		LogPath:            os.TempDir(),
	}
	for _, option := range opts {
		option(&op)
	}
	w := &Welfare{
		ctx:     ctx,
		channel: channel,
		gateway: gateway,
		logger:  log.New(ctx, op.LogPath, op.Loglevel),
		opt:     op,
	}

	if len(op.PrivateKey) > 0 && len(op.PublicKey) > 0 {
		w.secret = &rsa.RSASecret{
			PublicKey:          op.PublicKey,
			PrivateKey:         op.PrivateKey,
			PublicKeyDataType:  op.PublicKeyDataType,
			PrivateKeyType:     op.PrivateKeyType,
			PrivateKeyDataType: op.PrivateKeyDataType,
		}
	}

	return w
}

// InitSecret .
func (w *Welfare) InitSecret(ctx context.Context) error {
	if len(w.opt.PublicKey) == 0 || len(w.opt.PrivateKey) == 0 {
		w.logger.Log.Info(ErrInvalidPublicKeyOrPrivateKey.Error())
		return ErrInvalidPublicKeyOrPrivateKey
	}
	w.secret = &rsa.RSASecret{
		PublicKey:          w.opt.PublicKey,
		PrivateKey:         w.opt.PrivateKey,
		PublicKeyDataType:  w.opt.PublicKeyDataType,
		PrivateKeyType:     w.opt.PrivateKeyType,
		PrivateKeyDataType: w.opt.PrivateKeyDataType,
	}
	return nil
}

// Init .
func (w *Welfare) Init(ctx context.Context, opts ...Option) error {
	op := w.opt
	for _, option := range opts {
		option(&op)
	}
	if len(op.PublicKey) == 0 || len(op.PrivateKey) == 0 {
		w.logger.Log.Info(ErrInvalidPublicKeyOrPrivateKey.Error())
		return ErrInvalidPublicKeyOrPrivateKey
	}
	w.secret = &rsa.RSASecret{
		PublicKey:          op.PublicKey,
		PrivateKey:         op.PrivateKey,
		PublicKeyDataType:  op.PublicKeyDataType,
		PrivateKeyType:     op.PrivateKeyType,
		PrivateKeyDataType: op.PrivateKeyDataType,
	}
	return nil
}

// Context .
func (w *Welfare) Context() context.Context {
	return w.ctx
}

// Gateway .
func (w *Welfare) Gateway(ctx context.Context) string {
	return w.gateway
}

// Channel .
func (w *Welfare) Channel(ctx context.Context) string {
	return w.channel
}

// ServiceType .
func (w *Welfare) ServiceType(ctx context.Context) string {
	return w.serviceType
}

// Logger .
func (w *Welfare) Logger(ctx context.Context) *log.Logger {
	return w.logger
}

// Options .
func (w *Welfare) Options(ctx context.Context) Options {
	return w.opt
}

// SetServiceType .
func (w *Welfare) SetServiceType(ctx context.Context, serviceType string) {
	w.serviceType = serviceType
}

// SetLogger .
func (w *Welfare) SetLogger(ctx context.Context, logger *log.Logger) {
	w.logger = logger
}

// SetChannel .
func (w *Welfare) SetChannel(ctx context.Context, channel string) {
	w.channel = channel
}

// SetSecret .
func (w *Welfare) SetSecret(ctx context.Context, secret *rsa.RSASecret) {
	w.secret = secret
}

// SetContext .
func (w *Welfare) SetContext(ctx context.Context) {
	w.ctx = ctx
}

// SetGateway .
func (w *Welfare) SetGateway(ctx context.Context, gateway string) {
	w.gateway = gateway
}

// Secret .
func (w *Welfare) Secret(ctx context.Context) (rsa.RSASecret, error) {
	if len(w.opt.PublicKey) == 0 || len(w.opt.PrivateKey) == 0 {
		w.logger.Log.Info(ErrInvalidPublicKeyOrPrivateKey.Error())
		return rsa.RSASecret{}, ErrInvalidPublicKeyOrPrivateKey
	}
	return *w.secret, nil
}

// Close 关闭
func (w *Welfare) Close(ctx context.Context) error {
	if w.logger != nil {
		return w.logger.Log.Sync()
	}
	return nil
}

// Sync 同步
func (w *Welfare) Sync(ctx context.Context) error {
	if w.logger != nil {
		return w.logger.Log.Sync()
	}
	return nil
}
