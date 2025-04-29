package consts

//   `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态: 1-待审核(Pending), 2-处理中(Processing), 3-已拒绝(Rejected), 4-已完成(Completed), 5-失败(Failed)',

const (
	UserWithdrawStatusPending    = "1" // 待处理
	UserWithdrawStatusProcessing = "2" // 等待人工处理
	UserWithdrawStatusRejected   = "3" // 已拒绝
	UserWithdrawStatusCompleted  = "4" // 已完成
	UserWithdrawStatusFailed     = "5" // 失败
)
