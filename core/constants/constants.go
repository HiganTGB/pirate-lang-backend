package constants

import "time"

const (
	AccessTokenExpiry  = 24 * time.Hour
	RefreshTokenExpiry = 7 * 24 * time.Hour
)

// Limit login times
const (
	MaxLoginAttempts = 5
	BlockDuration    = 15 * time.Minute
)
