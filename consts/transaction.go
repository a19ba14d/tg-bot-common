package consts

// Transaction Types
type TransactionType string

const (
	TransactionTypeDeposit         TransactionType = "deposit"
	TransactionTypeWithdraw        TransactionType = "withdraw"          // 提现
	TransactionTypeTransferOut     TransactionType = "transfer_out"      // 转账支出
	TransactionTypeTransferIn      TransactionType = "transfer_in"       // 转账收入
	TransactionTypeRedPacketCreate TransactionType = "red_packet_create" // 创建红包（支出）
	TransactionTypeRedPacketClaim  TransactionType = "red_packet_claim"  // 领取红包（收入）
	TransactionTypeRedPacketRefund TransactionType = "red_packet_refund" // 红包退款（收入）
	// Add other transaction types as needed
)

// Transaction Directions
const (
	TransactionDirectionIn  = "in"  // 资金增加
	TransactionDirectionOut = "out" // 资金减少
)

// Entity Types (for related_entity_type in transactions)
const (
	EntityTypeWithdrawal = "withdrawal"
	EntityTypeRedPacket  = "red_packet"
	EntityTypeTransfer   = "transfer"
	// Add other entity types as needed
)
