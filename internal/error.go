package internal

import (
	"errors"
)

var (
	// ErrInvalidOptionLength .
	ErrInvalidOptionLength = errors.New("invalid option length")

	// ErrInvalidChannel .
	ErrInvalidChannel = errors.New("channel is empty")

	// ErrInvalidEnvironment .
	ErrInvalidEnvironment = errors.New("environment is empty")

	// ErrInvalidPublicKey .
	ErrInvalidPublicKey = errors.New("publicKey is empty")

	// ErrInvalidPrivateKey .
	ErrInvalidPrivateKey = errors.New("privateKey is empty")

	// ErrInvalidPublicKeyDataType .
	ErrInvalidPublicKeyDataType = errors.New("publicKeyDataType is empty")

	// ErrInvalidPrivateKeyType .
	ErrInvalidPrivateKeyType = errors.New("privateKeyType is empty")

	// ErrInvalidPrivateKeyDataType .
	ErrInvalidPrivateKeyDataType = errors.New("privateKeyDataType is empty")

	// ErrInvalidSignType .
	ErrInvalidSignType = errors.New("signType is empty")

	// ErrInvalidEncryptType .
	ErrInvalidEncryptType = errors.New("encryptType is empty")

	// ErrInvalidFormat .
	ErrInvalidFormat = errors.New("format is empty")

	// ErrInvalidVersion .
	ErrInvalidVersion = errors.New("version is empty")

	// ErrInvalidSign .
	ErrInvalidSign = errors.New("sign is empty")

	// ErrInvalidEncrypt .
	ErrInvalidEncrypt = errors.New("encrypt is empty")

	// ErrInvalidDecrypt .
	ErrInvalidDecrypt = errors.New("decrypt is empty")

	// ErrInvalidPublicKeyOrPrivateKey .
	ErrInvalidPublicKeyOrPrivateKey = errors.New("publicKey or privateKey is empty")
)
