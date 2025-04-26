package consts

// contextKey 是上下文键的自定义类型
type ContextKey string

const (
	ContextKeyBusinessType ContextKey = "business_type" // Context key for business type
	// KafkaTopicIncoming 定义接收来自 Telegram 网关消息的 Kafka Topic
	KafkaTopicIncoming = "telegram_incoming_messages"
	// KafkaTopicOutgoing 定义发送回复到 Telegram 网关的 Kafka Topic
	KafkaTopicOutgoing             = "telegram_outgoing_replies"
	KafkaTopicPaymentNotifications = "payment_notifications" // 收款成功通知 Topic
	// KafkaTopicDeadLetter 定义死信队列的 Kafka Topic (可选)
	// KafkaTopicDeadLetter = "telegram_dead_letter_queue"
)

// Business Types
const (
	// ... existing types ...
	BusinessTypeVerifyBackupAccount = "verify_backup_account" // 用于备用账户验证时的支付密码校验
	BusinessTypeTransferConfirm     = "transfer_confirm"      // 用于点对点转账确认时的支付密码校验
	BusinessTypeVerifyTransfer      = "verify_transfer"       // 用于 Start 参数触发的转账密码验证
	BusinessTypeRedPacket           = "redpacket"             // 用于红包创建时的支付密码校验
	BusinessTypePaymentRequest      = "payment_request"       // 用于收款请求支付时的支付密码校验 (新添加)
)

// Transfer Status
const (
	TransferStatusPending                 = "pending"            // 注意：此状态可能需要根据新状态调整或弃用
	TransferStatusPendingPass             = "pending_pass"       // 等待发送方密码验证
	TransferStatusPendingDeduction        = "pending_deduction"  // 等待系统扣款
	TransferStatusPendingCollection       = "pending_collection" // 等待收款人领取
	TransferStatusCompleted               = "completed"
	TransferStatusExpired                 = "expired"
	TransferStatusFailed                  = "failed"                    // 转账失败（通用）
	TransferStatusFailedInsufficientFunds = "failed_insufficient_funds" // 因余额不足导致失败
	TransferStatusCancelled               = "cancelled"                 // 转账被取消
)

// Payment Request Status
const (
	PaymentRequestStatusPending   = 1 // 待支付
	PaymentRequestStatusPaid      = 2 // 已支付
	PaymentRequestStatusExpired   = 3 // 已过期
	PaymentRequestStatusCancelled = 4 // 已取消
)

// User States for Inline Transfer

// Start Command Prefixes
const (
	StartParamPrefixPaymentRequest = "payreq_" // /start 命令处理收款请求的前缀 (e.g., /start payreq_USDT_6281127279)
)
const (
	StateInlineTransferAwaitingAmount = "inline_transfer_awaiting_amount" // 等待用户在特定聊天输入金额
	// StateInlineTransferAwaitingRecipient = "inline_transfer_awaiting_recipient" // (如果采用分步选择收款人)
	StateWaitingForTransferPassword = "waiting_transfer_password" // 等待用户输入转账支付密码
)

// Callback Prefixes
const (
	// ... existing prefixes ...
	CallbackPrefixVerifyBackupAccount  = "verify_backup_"     // 触发验证流程按钮
	CallbackPrefixVerifyBackupPassword = "verify_backup_pwd_" // 验证密码数字键盘
)

// Token Precision (adjust based on your system's needs, e.g., 8 for USDT/TRX)
const (
	TokenPrecision = 4
)

// Inline Query Prefixes
const (
	InlineQueryPrefixSearch     = "t:" // Prefix for recipient search query (transfer confirmation)
	InlineQueryPrefixPayRequest = "r:" // Prefix for payment request inline query
)
