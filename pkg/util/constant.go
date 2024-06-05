package util

import "time"

const (
	EXPIRED_CODE_DURATION      = time.Minute * 5
	LIMIT_RESEND_CODE_DURATION = time.Minute * 3
	JWT_DURATION               = time.Hour * 24
	JWT_USER_ID_CONTEXT_KEY    = "USER_ID"
)
