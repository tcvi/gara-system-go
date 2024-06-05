package util

import "time"

const (
	EXPIRED_CODE_DURATION      = time.Minute * 5
	LIMIT_RESEND_CODE_DURATION = time.Minute * 3
)
