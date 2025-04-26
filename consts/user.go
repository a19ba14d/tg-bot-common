package consts

// CtxKey defines the type for context keys.
type CtxKey string

const (
	// CtxUserKey is the context key for storing user information (*entity.Users).
	CtxUserKey CtxKey = "UserInfo"
)
