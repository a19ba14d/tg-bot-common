package consts

// User States for state machine
const (
	// ... other states ...

	// Red Packet Cover Upload State
	UserStateWaitingRedPacketCover = "waiting_rp_cover" // Waiting for user to send red packet cover image


	// Red Packet Creation States
	UserStateWaitingRedPacketQuantity = "waiting_rp_quantity" // Waiting for user to input red packet quantity
	// ... other states ...
	UserStateWaitingRedPacketAmount = "waiting_rp_amount" // Waiting for user to input red packet amount

	UserStateWaitingRedPacketBlessing = "waiting_rp_blessing" // 等待用户输入红包留言

	UserStateWaitingRedPacketCoverSelection = "waiting_rp_cover_selection" // 等待用户选择或确认红包封面
	UserStateWaitingRedPacketPaymentPassword = "waiting_rp_payment_password" // 等待用户输入红包支付密码
	StateWaitingForPaymentRequestPassword = "StateWaitingForPaymentRequestPassword" // 等待用户输入收款请求支付密码
)
// Input Types (used in UserState.ExpectedInputType)
const (
	InputTypeRedPacketPaymentPassword = "rp_payment_password" // Indicates expecting input via red packet password keyboard
	InputTypeText                     = "text"                // Default: expecting regular text input
	// Add other specific input types as needed, e.g., "photo", "contact"
)
const StateWaitingForReceivePayer = "waiting_for_receive_payer" // State indicating the bot is waiting for the user to select a payer for receiving payment