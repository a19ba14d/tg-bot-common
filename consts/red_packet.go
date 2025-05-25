package consts

// RedPacketImageStatus defines the possible statuses for red packet cover images.
type RedPacketImageStatus string

const (
	// RedPacketImageStatusPendingReview indicates the image is uploaded and waiting for review.
	RedPacketImageStatusPendingReview RedPacketImageStatus = "pending_review"
	// RedPacketImageStatusSuccess indicates the image has been approved.
	RedPacketImageStatusSuccess RedPacketImageStatus = "success"
	// RedPacketImageStatusFail indicates the image has been rejected.
	RedPacketImageStatusFail RedPacketImageStatus = "fail"

	// Context keys for UserState during red packet creation
	RpContextKeyTokenSymbol              = "rp_token_symbol"                 // Stores the selected token symbol
	RpContextKeyType                     = "rp_type"                         // "fixed" or "random"
	RpContextKeyQuantity                 = "rp_quantity"                     // Context key for quantity
	RpContextKeyAmount                   = "rp_amount"                       // Context key for amount
	RpContextKeyBlessing                 = "rp_blessing"                     // 红包留言
	RpContextKeySelectedCoverFileId      = "rp_selected_cover_file_id"       // 已选封面 FileID (可能来自自定义选择或默认)
	RpContextKeyCurrentCustomCoverFileId = "rp_current_custom_cover_file_id" // 当前显示的自定义封面 FileID (临时存储)
	RpContextKeyChatID                   = "rp_chat_id"                      // Stores the chat ID where the red packet interaction happens
	RpContextKeyTotalAmount              = "rp_total_amount"                 // Context key for the calculated total amount (used for pass-free check)
	RpContextKeyCurrentPassword          = "rp_current_password"             // Context key for storing the currently entered password digits
)

// IsValid checks if the status string is a valid RedPacketImageStatus.
func (s RedPacketImageStatus) IsValid() bool {
	switch s {
	case RedPacketImageStatusPendingReview, RedPacketImageStatusSuccess, RedPacketImageStatusFail:
		return true
	default:
		return false
	}
}

// RedPacketType defines the type of red packet distribution.
type RedPacketType string

const (
	RedPacketTypeRandom RedPacketType = "random" // Random amount per claim
	RedPacketTypeFixed  RedPacketType = "fixed"  // Fixed amount per claim
)

// IsValid checks if the type string is a valid RedPacketType.
func (t RedPacketType) IsValid() bool {
	switch t {
	case RedPacketTypeRandom, RedPacketTypeFixed:
		return true
	default:
		return false
	}
}

// RedPacketStatus defines the overall status of a red packet.
type RedPacketStatus string

const (
	RedPacketStatusActive    RedPacketStatus = "active"    // Red packet is active and can be claimed. (Matches DB 'active')
	RedPacketStatusEmpty     RedPacketStatus = "empty"     // All claims have been made.
	RedPacketStatusExpired   RedPacketStatus = "expired"   // Red packet has passed its expiration time.
	RedPacketStatusCancelled RedPacketStatus = "cancelled" // Red packet has been cancelled by the creator.
)

// RedPacketClaimStatus defines the status of an individual claim record.
type RedPacketClaimStatus string

const (
	RedPacketClaimStatusPending   RedPacketClaimStatus = "pending"   // The claim record is available but not yet claimed.
	RedPacketClaimStatusClaimed   RedPacketClaimStatus = "claimed"   // The claim has been successfully made by a user.
	RedPacketClaimStatusCancelled RedPacketClaimStatus = "cancelled" // The claim has been cancelled (e.g., due to red packet cancellation).
)

const (
	// Wallet business type for red packet creation debit
	WalletBusinessTypeRedPacketCreate = "red_packet_create"
	// Wallet business type for red packet claim credit
	WalletBusinessTypeRedPacketClaim = "red_packet_claim"
	// Wallet business type for expired red packet refund credit
	WalletBusinessTypeRedPacketRefund = "red_packet_refund"
	// Wallet business type for cancelled red packet refund credit
	WalletBusinessTypeRedPacketCancelRefund = "red_packet_cancel_refund"

	// Minimum amount allowed for a single claim (adjust as needed)
	RedPacketMinClaimAmount = 0.001

	// Maximum number of red packets allowed per creation
	RedPacketMaxQuantity = 100
)
