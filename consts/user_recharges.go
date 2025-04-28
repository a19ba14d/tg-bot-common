package consts

//  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态: 1-待确认/处理中(Pending), 2-已完成/已入账(Completed)',

const (
	UserRechargeStatusPending = "1" // 待确认/处理中
	UserRechargeStatusSuccess = "2" // 已完成/已入账
	UserRechargeStatusFailed  = "3" // 失败
)
