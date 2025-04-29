package codes

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 通用错误码
var (
	// 成功
	CodeSuccess = gcode.New(0, "操作成功", nil)
	// 通用错误
	CodeError = gcode.New(1, "操作失败", nil)
	// 参数错误
	CodeInvalidParameter = gcode.New(2, "参数错误", nil)
	// 未授权
	CodeUnauthorized = gcode.New(401, "未授权或授权已过期", nil)
	// 无权限
	CodeForbidden = gcode.New(403, "无权限访问", nil)
	// 资源不存在
	CodeNotFound = gcode.New(404, "资源不存在", nil)
	// 服务器内部错误
	CodeInternalError = gcode.New(500, "服务器内部错误", nil)

	// 业务错误码 - 认证模块 (1000-1999)
	CodeLoginFailed    = gcode.New(1001, "登录失败，账号或密码错误", nil)
	CodeInvalidCaptcha = gcode.New(1002, "验证码错误", nil)
	CodeAccountLocked  = gcode.New(1003, "账号已被锁定", nil)
	CodeTokenExpired   = gcode.New(1004, "令牌已过期", nil)
	CodeInvalidToken   = gcode.New(1005, "无效的令牌", nil)

	// 通用 - 功能未实现
	CodeNotImplemented = gcode.New(501, "功能暂未实现", nil)

	// 业务错误码 - Exchange 模块 (3000-3999)
	CodeSystemExchangeProductNotFound         = gcode.New(3001, "兑换产品不存在", nil)
	CodeSystemExchangeProductSymbolExists     = gcode.New(3002, "兑换产品交易对符号已存在", nil)
	CodeSystemExchangeOrderNotFound           = gcode.New(3011, "兑换订单不存在", nil)
	CodeSystemExchangeOrderCannotUpdateStatus = gcode.New(3012, "当前订单状态无法修改", nil)
	CodeSystemExchangeProductPairExists       = gcode.New(3003, "兑换产品基础/计价代币对已存在", nil)

	// 业务错误码 - 商户模块 (4000-4999)
	CodeMerchantNotFound     = gcode.New(4001, "商户不存在", nil)
	CodeMerchantNameExists   = gcode.New(4002, "商户名称已存在", nil)
	CodeMerchantUserIdExists = gcode.New(4003, "该用户已关联商户", nil)
	CodeApiKeyNotFound       = gcode.New(4101, "API密钥不存在", nil)
	CodeApiKeyRevoked        = gcode.New(4102, "API密钥已被撤销", nil)
	CodeApiKeyExpired        = gcode.New(4103, "API密钥已过期", nil)

	// 业务错误码 - 用户模块 (5000-5999)
	CodeUserNotFound             = gcode.New(5001, "用户不存在", nil)
	CodeUserAccountExists        = gcode.New(5002, "账号已存在", nil)
	CodeUserEmailExists          = gcode.New(5003, "邮箱已存在", nil)
	CodeUserPhoneExists          = gcode.New(5004, "手机号已存在", nil)
	CodeUserInviteCodeExists     = gcode.New(5005, "邀请码已存在", nil)
	CodeBackupAccountNotFound    = gcode.New(5101, "备用账户不存在", nil)
	CodeBackupAccountExists      = gcode.New(5102, "备用账户已存在", nil)
	CodeInvalidVerificationToken = gcode.New(5103, "验证令牌无效", nil)
	CodeVerificationTokenExpired = gcode.New(5104, "验证令牌已过期", nil)

	// 业务错误码 - 钱包/交易模块 (2000-2999)
	CodeInsufficientBalance = gcode.New(2001, "账户余额不足", nil)
	CodeWalletError         = gcode.New(2002, "钱包服务错误", nil)
	CodeDatabaseError       = gcode.New(2003, "数据库操作失败", nil)
	CodeValidationFailed    = gcode.New(2004, "参数验证失败", nil)

	// 业务错误码 - 支付密码相关 (5200-5299)
	CodePasswordIncorrect = gcode.New(5201, "支付密码错误", nil)
	CodePasswordNotSet    = gcode.New(5202, "未设置支付密码", nil)

	// 业务错误码 - 支付密码相关 (5200-5299)
	CodePasswordLocked = gcode.New(5203, "支付密码被锁定", nil)

	// 业务错误码 - 转账模块 (8000-8099)
	CodeTransferFailed            = gcode.New(8000, "转账处理失败", nil)
	CodeTransferNotReceiver       = gcode.New(8001, "非指定收款人", nil)
	CodeTransferIsSender          = gcode.New(8002, "不能领取自己发起的转账", nil)
	CodeTransferAlreadyCompleted  = gcode.New(8003, "转账已被领取", nil)
	CodeTransferExpired           = gcode.New(8004, "转账已过期", nil)
	CodeTransferInsufficientFunds = gcode.New(8005, "发送者余额不足", nil)
	CodeTransferPendingPass       = gcode.New(8006, "转账等待付款方验证", nil)
	CodeTransferInvalidStatus     = gcode.New(8007, "转账状态无效", nil)

	// 业务错误码 - 收款请求模块 (9000-9099)
	CodePaymentRequestNotFound      = gcode.New(9001, "收款请求不存在或已处理", nil)
	CodePaymentRequestAlreadyPaid   = gcode.New(9002, "该收款请求已支付", nil)
	CodePaymentRequestExpired       = gcode.New(9003, "该收款请求已过期", nil)
	CodePaymentRequestCancelled     = gcode.New(9004, "该收款请求已取消", nil)
	CodePaymentRequestCannotPaySelf = gcode.New(9005, "不能向自己支付", nil)
	CodePaymentRequestInvalidStatus = gcode.New(9006, "无效的收款请求状态", nil)

	// 业务错误码 - 日志模块 (6000-6999)
	CodeOperationLogNotFound = gcode.New(6001, "操作日志不存在", nil)
	// 业务错误码 - 内部逻辑 (7000-7999)
	CodeExecutorNotFound = gcode.New(7001, "未找到对应的回调执行器", nil)
)

// Sentinel error to indicate a message was handled internally by a state handler
// and the registry should stop processing further handlers.
var ErrMessageHandled = gerror.NewSkip(1, "message handled internally")

// Removed misplaced code

// Removed incorrectly placed code definition

// NewError 创建一个带错误码的错误，对于业务错误不记录堆栈信息
func NewError(code gcode.Code) error {
	// 业务错误码范围 1000-9999，不记录堆栈
	if code.Code() >= 1000 && code.Code() <= 9999 {
		// 跳过堆栈收集
		return gerror.NewCodeSkip(code, 1)
	}
	// 系统错误记录完整堆栈
	return gerror.NewCode(code)
}

// NewErrorf 创建一个带错误码和格式化信息的错误
func NewErrorf(code gcode.Code, format string, args ...interface{}) error {
	// 业务错误码范围 1000-9999，不记录堆栈
	if code.Code() >= 1000 && code.Code() <= 9999 {
		// 跳过堆栈收集，但保留格式化消息
		skip := 1
		return gerror.NewCodeSkipf(code, skip, format, args...)
	}
	// 系统错误记录完整堆栈
	return gerror.NewCodef(code, format, args...)
}

// NewErrorSkip 创建一个带错误码和调用栈跳过的错误
func NewErrorSkip(code gcode.Code, skip int) error {
	return gerror.NewCodeSkip(code, skip)
}

// WrapError 包装一个错误并添加错误码
func WrapError(err error, code gcode.Code) error {
	// 业务错误码范围 1000-9999，不记录额外的堆栈
	if code.Code() >= 1000 && code.Code() <= 9999 {
		// 只记录错误信息，不记录堆栈
		return gerror.NewCodeSkipf(code, 1, "%v", err.Error())
	}
	// 系统错误记录完整堆栈
	return gerror.WrapCode(code, err)
}
