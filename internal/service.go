package internal

import (
	"context"
)

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
