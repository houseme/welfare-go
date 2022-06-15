package internal

import (
	"context"
	"time"

	"github.com/houseme/gocrypto"

	"github.com/houseme/welfare-go/pkg/log"
)

const (
	// devGateway .测试环境网关地址
	devGateway = "https://sandbox-api-welfare.wasair.com/api.v2/partner/base"

	// prodGateway .生产环境网关地址
	prodGateway = "https://api-welfare.wasair.com/api.v2/partner/base"

	// version .
	version = "2.0.0"

	// devEnvironment .
	devEnvironment = "dev"

	// prodEnvironment .
	prodEnvironment = "prod"

	format = "json"

	// encryptType Encrypt Type
	encryptType = "RSA2"

	// signType sign type
	signType = "RSA2"

	// publicKeyDataType public key data type
	publicKeyDataType = gocrypto.Base64

	// privateKeyType private key type
	privateKeyType = gocrypto.PKCS1

	// privateKeyDataType private key data type
	privateKeyDataType = gocrypto.Base64

	// timeout time out 5*time.Second
	timeout = 5 * time.Second
)

// Option The option is a polaris option.
type Option func(o *Options)

type Options struct {
	// Version
	Version string `json:"version"`

	// required, Format access token
	Format string

	// optional, EncryptType in welfare. Default value is RSA2
	EncryptType string

	// service sign type in welfare. Default value is RSA2
	SignType string

	// service public key. Default value is ""
	PublicKey string

	// To show the public key data type. Default value is gocrypto.Base64.
	PublicKeyDataType gocrypto.Encode

	// PrivateKey in welfare . Default value is True.
	PrivateKey string

	// To show the private key data type. Default value is gocrypto.Base64.
	PrivateKeyDataType gocrypto.Encode

	// PrivateKeyType in welfare. Default value is gocrypto.Secret.
	PrivateKeyType gocrypto.Secret

	// optional, Timeout for the single query. Default value is global config
	// Total is (5 time.Second) * Timeout
	Timeout time.Duration

	// optional, retry count. Default value is global config
	RetryCount int

	// optional, LogPath. Default value is os.TempDir()
	LogPath string `json:"logPath"`

	// optional, Loglevel. Default value is log.Info
	Loglevel log.Level `json:"logLevel"`
}

// SetPublicKey .
func (w *Welfare) SetPublicKey(ctx context.Context, publicKey string) {
	w.opt.PublicKey = publicKey
}

// SetPrivateKey .
func (w *Welfare) SetPrivateKey(ctx context.Context, privateKey string) {
	w.opt.PrivateKey = privateKey
}

// SetPublicKeyDataType .
func (w *Welfare) SetPublicKeyDataType(ctx context.Context, publicKeyDataType gocrypto.Encode) {
	w.opt.PublicKeyDataType = publicKeyDataType
}

// SetPrivateKeyType .
func (w *Welfare) SetPrivateKeyType(ctx context.Context, privateKeyType gocrypto.Secret) {
	w.opt.PrivateKeyType = privateKeyType
}

// SetPrivateKeyDataType .
func (w *Welfare) SetPrivateKeyDataType(ctx context.Context, privateKeyDataType gocrypto.Encode) {
	w.opt.PrivateKeyDataType = privateKeyDataType
}

// SetPublicKeyAndPrivateKey  .
func (w *Welfare) SetPublicKeyAndPrivateKey(ctx context.Context, publicKey, privateKey string) {
	w.opt.PublicKey = publicKey
	w.opt.PrivateKey = privateKey
}

// SetTimeout .
func (w *Welfare) SetTimeout(ctx context.Context, timeout time.Duration) {
	w.opt.Timeout = timeout
}

// SetRetryCount .
func (w *Welfare) SetRetryCount(ctx context.Context, retryCount int) {
	w.opt.RetryCount = retryCount
}

// SetLogPath .
func (w *Welfare) SetLogPath(ctx context.Context, logPath string) {
	w.opt.LogPath = logPath
}

// SetLoglevel .
func (w *Welfare) SetLoglevel(ctx context.Context, loglevel log.Level) {
	w.opt.Loglevel = loglevel
}

// SetVersion .
func (w *Welfare) SetVersion(ctx context.Context, version string) {
	w.opt.Version = version
}

// SetFormat .
func (w *Welfare) SetFormat(ctx context.Context, format string) {
	w.opt.Format = format
}

// SetEncryptType .
func (w *Welfare) SetEncryptType(ctx context.Context, encryptType string) {
	w.opt.EncryptType = encryptType
}

// SetSignType .
func (w *Welfare) SetSignType(ctx context.Context, signType string) {
	w.opt.SignType = signType
}

// version .
func (w *Welfare) Version(ctx context.Context) string {
	return w.opt.Version
}

// format .
func (w *Welfare) Format(ctx context.Context) string {
	return w.opt.Format
}

// signType .
func (w *Welfare) SignType(ctx context.Context) string {
	return w.opt.SignType
}

// encryptType .
func (w *Welfare) EncryptType(ctx context.Context) string {
	return w.opt.EncryptType
}

// RetryCount .
func (w *Welfare) RetryCount(ctx context.Context) int {
	return w.opt.RetryCount
}

// PublicKey .
func (w *Welfare) PublicKey(ctx context.Context) string {
	return w.opt.PublicKey
}

// PrivateKey .
func (w *Welfare) PrivateKey(ctx context.Context) string {
	return w.opt.PrivateKey
}

// publicKeyDataType .
func (w *Welfare) PublicKeyDataType(ctx context.Context) gocrypto.Encode {
	return w.opt.PublicKeyDataType
}

// privateKeyType .
func (w *Welfare) PrivateKeyType(ctx context.Context) gocrypto.Secret {
	return w.opt.PrivateKeyType
}

// privateKeyDataType .
func (w *Welfare) PrivateKeyDataType(ctx context.Context) gocrypto.Encode {
	return w.opt.PrivateKeyDataType
}

// LogPath .
func (w *Welfare) LogPath(ctx context.Context) string {
	return w.opt.LogPath
}

// Loglevel .
func (w *Welfare) Loglevel(ctx context.Context) log.Level {
	return w.opt.Loglevel
}

// timeout .
func (w *Welfare) Timeout(ctx context.Context) time.Duration {
	return w.opt.Timeout
}
