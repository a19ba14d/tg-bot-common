package consts

// Callback Data Prefixes for Withdraw Flow
const (
	CallbackActionNoOp          = "noop"      // Placeholder for buttons that do nothing
	WithdrawStart               = "withdraw"  // Callback data to start the flow
	WithdrawPrefix              = "withdraw_" // General prefix for the flow
	WithdrawSelectSymbol        = WithdrawPrefix + "symbol_"
	WithdrawSelectChain         = WithdrawPrefix + "chain_"
	WithdrawSelectAddressType   = WithdrawPrefix + "addr_type_"
	WithdrawSelectWhitelistAddr = WithdrawPrefix + "wl_addr_"
	WithdrawAddNewWhitelistAddr = WithdrawPrefix + "add_wl_"
	WithdrawDeleteWhitelistAddr = WithdrawPrefix + "del_wl_" // Placeholder for delete action
	WithdrawConfirm             = WithdrawPrefix + "confirm"
	WithdrawCancel              = WithdrawPrefix + "cancel"
	WithdrawBackToSymbolSelect  = WithdrawPrefix + "back_symbol"
	WithdrawBackToChainSelect   = WithdrawPrefix + "back_chain_"
	WithdrawBackToAddressSelect = WithdrawPrefix + "back_addr_"
	WithdrawBackToAmountInput   = WithdrawPrefix + "back_amount"
	// Password keyboard callback prefixes
	WithdrawPasswordKeyboard        = WithdrawPrefix               // Prefix for password keyboard callbacks
	WithdrawAddressPasswordKeyboard = WithdrawPrefix + "addr_pwd_" // Prefix for address verification password keyboard

	// Withdraw History callback prefix
	WithdrawHistoryPrefix = "wh:" // Prefix for withdraw history callbacks
)

// Callback Action Prefix for Transfer
const (
	CallbackActionTransferCollect        = "transfer_collect:"    // 收款操作
	CallbackVerifyTransferPasswordPrefix = "verify_transfer_pwd:" // Prefix for verifying transfer password via callback
	StartParamVerifyTransferPrefix       = "verify_transfer_"     // Prefix for verifying transfer via start parameter
	CallbackPrefixVerifyTransferPassword = "verify_transfer_pwd" // 转账密码验证数字键盘回调前缀 (Removed trailing underscore)
)

// Inline Query Prefixes/Types
const (
	InlineQueryTypeTransferInitiate = "transfer_init" // 发起转账的内联查询标识
	InlineQueryPrefixShareRedPacket = "share_rp "     // Prefix for sharing red packet via inline query
	// (可能需要其他类型用于区分内联查询结果的后续操作)
	CallbackPrefixClaimRedPacket = "claim_rp_" // Prefix for claiming red packet via callback

	// Transfer Callbacks
	CallbackActionTransferPassword = "transfer_pwd:"     // 支付密码键盘
	CallbackActionTransferConfirm  = "transfer_confirm:" // 确认转账（免密）
	CallbackActionTransferCancel   = "transfer_cancel"   // 取消转账

	InlineCallbackDataAmountPrompt = "transfer_amount_prompt:" // (如果使用CallbackData传递上下文)
)

// Add other callback prefixes here if needed, e.g., for deposit, profile, etc.

// Callback Prefixes/Data for Navigation
const (
	CallbackNavigateBack = "navigation:main" // General back navigation
)

// Callback Prefixes/Data for Profile and PassFree Settings
const (
	CallbackProfile = "profile" // Callback data to show the main profile view

	CallbackPassFreeSetLimitPrefix = "set_limit_" // Prefix for setting pass-free limit for a symbol

	CallbackRedPacket = "redpacket" // Main menu button callback

	// Callback Prefixes/Data for Red Packet Flow
	CallbackRedPacketPrefix   = "rp:" // General prefix for the flow
	CallbackRedPacketOngoing  = CallbackRedPacketPrefix + "ongoing"
	CallbackRedPacketEnded    = CallbackRedPacketPrefix + "ended"
	CallbackRedPacketAdd      = CallbackRedPacketPrefix + "add"
	CallbackRedPacketSetCover = CallbackRedPacketPrefix + "set_cover"

	// New constants for token and type selection
	CallbackSelectRedPacketTokenPrefix = CallbackRedPacketPrefix + "select_symbol_" // Prefix for selecting token: rp:select_symbol_{symbol}
	CallbackSelectRedPacketTypeRandom  = CallbackRedPacketPrefix + "type_random"    // Select random type: rp:type_random:{symbol}
	CallbackSelectRedPacketTypeFixed   = CallbackRedPacketPrefix + "type_fixed"     // Select fixed type: rp:type_fixed:{symbol}
	CallbackRedPacketBackToTokenSelect = CallbackRedPacketPrefix + "back_token"     // Back to token selection

	// Red Packet Cover specific callbacks
	CallbackRedPacketCoverPagePrefix       = CallbackRedPacketPrefix + "cover_page_"         // 分页: rp:cover_page_{page}
	CallbackRedPacketCoverDelConfirmPrefix = CallbackRedPacketPrefix + "cover_del_confirm_"  // 删除确认: rp:cover_del_confirm_{imageId}_{currentPage}
	CallbackRedPacketCoverDelExecPrefix    = CallbackRedPacketPrefix + "cover_del_exec_"     // 执行删除: rp:cover_del_exec_{imageId}
	CallbackRedPacketCoverUploadPrompt     = CallbackRedPacketPrefix + "cover_upload_prompt" // 上传提示: rp:cover_upload_prompt

	// Red Packet Blessing specific callbacks
	CallbackRedPacketInputBlessing   = CallbackRedPacketPrefix + "input_blessing"   // 用户点击“输入留言”
	CallbackRedPacketSkipBlessing    = CallbackRedPacketPrefix + "skip_blessing"    // 用户点击“跳过”
	CallbackRedPacketConfirmBlessing = CallbackRedPacketPrefix + "confirm_blessing" // 用户点击“确认”留言
	CallbackRedPacketReEnterBlessing = CallbackRedPacketPrefix + "reenter_blessing" // 用户点击“重新输入”留言

	// Red Packet Cover Selection specific callbacks
	CallbackRedPacketSelectCoverConfirm             = CallbackRedPacketPrefix + "cover:confirm"        // 确认使用当前封面: rp:cover:confirm
	CallbackRedPacketSelectCoverCustom              = CallbackRedPacketPrefix + "cover:custom"         // 选择自定义封面: rp:cover:custom
	CallbackRedPacketSelectCoverCustomPagePrefix    = CallbackRedPacketPrefix + "cover:page:"          // 自定义封面分页: rp:cover:page:{page}
	CallbackRedPacketSelectCoverCustomSelectPrefix  = CallbackRedPacketPrefix + "cover:select:"        // 选择自定义封面: rp:cover:select:{fileId}
	CallbackRedPacketSelectCoverCustomSelectCurrent = CallbackRedPacketPrefix + "cover:select:current" // 选择当前显示的自定义封面: rp:cover:select:current
	CallbackRedPacketSelectCoverCustomBack          = CallbackRedPacketPrefix + "cover:back"           // 从自定义返回默认: rp:cover:back
)

// Red Packet Final Confirmation Callbacks
const (
	CallbackPrefixRedPacketFinalConfirm = "rp:final:confirm"
	CallbackPrefixRedPacketFinalCancel  = "rp:final:cancel"
)

// Red Packet Payment Password Keyboard Callbacks
const (
	CallbackRedPacketPasswordKeyPrefix = "rp:pwd:key:"                                  // 数字键: rp:pwd:key:{digit}
	CallbackRedPacketPasswordConfirm   = CallbackRedPacketPasswordKeyPrefix + "confirm" // 确认键
	CallbackRedPacketPasswordDelete    = CallbackRedPacketPasswordKeyPrefix + "delete"  // 删除键
	CallbackRedPacketPasswordCancel    = CallbackRedPacketPasswordKeyPrefix + "cancel"  // 取消键
)

// Callback Prefixes for Receive Payment Flow
const (
	CallbackPrefixReceiveSelectToken = "receive_select_token:" // 选择收款代币
	CallbackPrefixSharePayReq        = "share_pay_req:"        // 获取收款分享链接 (New)
	CallbackPrefixPayRequestPassword = "pay_req_pwd:"          // 收款请求密码键盘
	CallbackPrefixPayRequestConfirm  = "pay_req_confirm:"      // 收款请求免密确认
)

// Callback Prefixes/Data for Account Statement
const (
	CallbackAccountStatement       = "account_statement"        // Exact match for profile menu button
	CallbackAccountStatementPrefix = "as_"                      // Prefix for account statement flow (select, page)
)

