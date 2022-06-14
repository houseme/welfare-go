package internal

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/houseme/gocrypto"
	"github.com/houseme/gocrypto/rsa"

	"github.com/houseme/welfare-go/pkg"
)

const (
	// DevGateway .测试环境网关地址
	DevGateway = "https://sandbox-api-welfare.wasair.com/api.v2/partner/base"

	// ProdGateway .生产环境网关地址
	ProdGateway = "https://api-welfare.wasair.com/api.v2/partner/base"

	// Version .
	Version = "2.0.0"

	// DevEnvironment .
	DevEnvironment = "dev"

	// ProdEnvironment .
	ProdEnvironment = "prod"

	Format = "json"

	// EncryptType Encrypt Type
	EncryptType = "RSA2"

	// SignType sign type
	SignType = "RSA2"

	// PublicKeyDataType public key data type
	PublicKeyDataType = gocrypto.Base64

	// PrivateKeyType private key type
	PrivateKeyType = gocrypto.PKCS1

	// PrivateKeyDataType private key data type
	PrivateKeyDataType = gocrypto.Base64

	// Timeout time out 5*time.Second
	Timeout = 5 * time.Second
)

var (
	once   sync.Once
	secret rsa.RSASecret
)

// Welfare .
type Welfare struct {
	Gateway     string
	Channel     string
	ServiceType string
	Logger      *pkg.Logger
	Opt         Options
	Secret      *rsa.RSASecret
}

// Option The option is a polaris option.
type Option func(o *Options)

type Options struct {
	// Version
	Version string `json:"version"`

	// required, Channel in welfare
	Channel string

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
}

// 合作伙伴请求平台的 serviceType
const (
	// LrUniEntUserCreate 注册用户信息
	lrUniEntUserCreate = "lr.uni.ent.user.create"

	// LrUniEntUserAuthorization 注册用户授权
	lrUniEntUserAuthorization = "lr.uni.ent.user.authorization"

	// LrUniEntBalanceQry 用户账户余额查询
	lrUniEntBalanceQry = "lr.uni.ent.user.balance.qry"

	// LrUniEntOrderCreate 统一下单
	lrUniEntOrderCreate = "lr.uni.ent.order.create"

	// LrUniEntOrderQry 订单查询
	lrUniEntOrderQry = "lr.uni.ent.order.qry"
)

// 针对平台方请求第三方合作伙伴的 serviceType
const (
	// ThirdUniEntBalanceQry 用户账户余额查询
	thirdUniEntBalanceQry = "third.uni.ent.user.balance.qry"

	// ThirdUniEntDeduct 余额扣除
	thirdUniEntDeduct = "third.uni.ent.user.balance.deduct"

	// ThirdUniEntDeductStatus 余额扣除状态查询
	thirdUniEntDeductStatus = "third.uni.ent.user.balance.deduct.status"

	// ThirdUniEntRefund 余额退款
	thirdUniEntRefund = "third.uni.ent.user.balance.refund"

	// ThirdUniEntReFundQry 退款查询
	thirdUniEntReFundQry = "third.uni.ent.refund.qry"

	// ThirdUniEntCancel 订单撤销
	thirdUniEntCancel = "third.uni.ent.cancel"
)

// SetServiceType .
func (w *Welfare) SetServiceType(ctx context.Context, serviceType string) {
	w.ServiceType = serviceType
}

// SetLogger .
func (w *Welfare) SetLogger(ctx context.Context, logger *pkg.Logger) {
	w.Logger = logger
}

// SetChannel .
func (w *Welfare) SetChannel(ctx context.Context, channel string) {
	w.Channel = channel
}

// LrUniEntUserCreate 注册用户信息
func (w *Welfare) LrUniEntUserCreate(ctx context.Context) string {
	return lrUniEntUserCreate
}

// LrUniEntUserAuthorization 注册用户授权
func (w *Welfare) LrUniEntUserAuthorization(ctx context.Context) string {
	return lrUniEntUserAuthorization
}

// LrUniEntBalanceQry 用户账户余额查询
func (w *Welfare) LrUniEntBalanceQry(ctx context.Context) string {
	return lrUniEntBalanceQry
}

// LrUniEntOrderCreate 统一下单
func (w *Welfare) LrUniEntOrderCreate(ctx context.Context) string {
	return lrUniEntOrderCreate
}

// LrUniEntOrderQry 订单查询
func (w *Welfare) LrUniEntOrderQry(ctx context.Context) string {
	return lrUniEntOrderQry
}

// ThirdUniEntBalanceQry 用户账户余额查询
func (w *Welfare) ThirdUniEntBalanceQry(ctx context.Context) string {
	return thirdUniEntBalanceQry
}

// ThirdUniEntReFundQry 退款查询
func (w *Welfare) ThirdUniEntReFundQry(ctx context.Context) string {
	return thirdUniEntReFundQry
}

// ThirdUniEntRefund 退款申请
func (w *Welfare) ThirdUniEntRefund(ctx context.Context) string {
	return thirdUniEntRefund
}

// ThirdUniEntDeduct 余额扣除
func (w *Welfare) ThirdUniEntDeduct(ctx context.Context) string {
	return thirdUniEntDeduct
}

// ThirdUniEntDeductStatus 余额扣除查询
func (w *Welfare) ThirdUniEntDeductStatus(ctx context.Context) string {
	return thirdUniEntDeductStatus
}

// ThirdUniEntCancel 订单撤销
func (w *Welfare) ThirdUniEntCancel(ctx context.Context) string {
	return thirdUniEntCancel
}

// GetGateway .
func (w *Welfare) GetGateway(ctx context.Context) string {
	return w.Gateway
}

// GetChannel .
func (w *Welfare) GetChannel(ctx context.Context) string {
	return w.Channel
}

// GetServiceType .
func (w *Welfare) GetServiceType(ctx context.Context) string {
	return w.ServiceType
}

// GetLogger .
func (w *Welfare) GetLogger(ctx context.Context) *pkg.Logger {
	return w.Logger
}

// GetTimeout .
func (w *Welfare) GetTimeout(ctx context.Context) time.Duration {
	return w.Opt.Timeout
}

// SetSecret .
func (w *Welfare) SetSecret(ctx context.Context, secret *rsa.RSASecret) {
	w.Secret = secret
}

// InitRSA .
func (w *Welfare) InitRSA(ctx context.Context) error {
	if len(w.Opt.PublicKey) == 0 || len(w.Opt.PrivateKey) == 0 {
		w.Logger.Log.Error("publicKey or privateKey is empty")
		return errors.New("publicKey or privateKey is empty")
	}
	w.Secret = &rsa.RSASecret{
		PublicKey:          w.Opt.PublicKey,
		PrivateKey:         w.Opt.PrivateKey,
		PublicKeyDataType:  w.Opt.PublicKeyDataType,
		PrivateKeyType:     w.Opt.PrivateKeyType,
		PrivateKeyDataType: w.Opt.PrivateKeyDataType,
	}
	return nil
}

// GetSecret .
func (w *Welfare) GetSecret(ctx context.Context) rsa.RSASecret {
	return *w.Secret
}

// SetPublicKey .
func (w *Welfare) SetPublicKey(ctx context.Context, publicKey string) {
	w.Opt.PublicKey = publicKey
}

// SetPrivateKey .
func (w *Welfare) SetPrivateKey(ctx context.Context, privateKey string) {
	w.Opt.PrivateKey = privateKey
}

// SetPublicKeyDataType .
func (w *Welfare) SetPublicKeyDataType(ctx context.Context, publicKeyDataType gocrypto.Encode) {
	w.Opt.PublicKeyDataType = publicKeyDataType
}

// SetPrivateKeyType .
func (w *Welfare) SetPrivateKeyType(ctx context.Context, privateKeyType gocrypto.Secret) {
	w.Opt.PrivateKeyType = privateKeyType
}

// SetPrivateKeyDataType .
func (w *Welfare) SetPrivateKeyDataType(ctx context.Context, privateKeyDataType gocrypto.Encode) {
	w.Opt.PrivateKeyDataType = privateKeyDataType
}

// SetPublicKeyAndPrivateKey  .
func (w *Welfare) SetPublicKeyAndPrivateKey(ctx context.Context, publicKey, privateKey string) {
	w.Opt.PublicKey = publicKey
	w.Opt.PrivateKey = privateKey
}
