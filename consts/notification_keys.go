package consts

// Notification Keys for the Unified Notification System
const (
	NotificationKeyPaymentSent           = "payment_sent"     // For payer
	NotificationKeyPaymentReceived       = "payment_received" // For requester/receiver
	NotificationKeyDepositSuccess        = "deposit_success"
	NotificationKeyTransferExpired       = "transfer_expired"
	NotificationKeyRedPacketExpired      = "red_packet_expired"
	NotificationKeyRedPacketClaimedFully = "red_packet_claimed_fully" // Corresponds to RedPacketClaimedPayload
	NotificationKeyWithdrawalSuccess     = "withdrawal_success"
	NotificationKeyWithdrawalFailed      = "withdrawal_failed"
	// Add other notification keys as needed in the future
)
