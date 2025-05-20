package errors

type ErrorCode int

const (
	// ErrInternal System errors (6000-6999)
	ErrInternal      ErrorCode = 6000
	ErrConfiguration ErrorCode = 6001
	ErrThirdParty    ErrorCode = 6002
	ErrNetwork       ErrorCode = 6003
)
