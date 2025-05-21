package errors

type ErrorCode int

const (

	// Resource errors (3000-3999)
	ErrNotFound        ErrorCode = 3000
	ErrAlreadyExists   ErrorCode = 3001
	ErrResourceLocked  ErrorCode = 3002
	ErrResourceExpired ErrorCode = 3003
	// ErrInternal System errors (6000-6999)
	ErrInternal      ErrorCode = 6000
	ErrConfiguration ErrorCode = 6001
	ErrThirdParty    ErrorCode = 6002
	ErrNetwork       ErrorCode = 6003
)
