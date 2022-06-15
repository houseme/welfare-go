package welfare

import (
	"time"

	"github.com/houseme/gocrypto"

	"github.com/houseme/welfare-go/internal"
	"github.com/houseme/welfare-go/pkg/log"
)

// WithVersion .
func WithVersion(version string) internal.Option {
	return func(o *internal.Options) {
		o.Version = version
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

// WithLogPath .
func WithLogPath(logPath string) internal.Option {
	return func(o *internal.Options) {
		o.LogPath = logPath
	}
}

// WithLoglevel .
func WithLoglevel(loglevel log.Level) internal.Option {
	return func(o *internal.Options) {
		o.Loglevel = loglevel
	}
}
