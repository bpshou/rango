package tools

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func CreateTotpCode(secret string, period uint) (string, error) {
	var option = totp.ValidateOpts{
		// Number of seconds a TOTP hash is valid for. Defaults to 30 seconds.
		Period: period,
		// Periods before or after the current time to allow.  Value of 1 allows up to Period
		// of either side of the specified time.  Defaults to 0 allowed skews.  Values greater
		// than 1 are likely sketchy.
		Skew: 1,
		// Digits as part of the input. Defaults to 6.
		Digits: 6,
		// Algorithm to use for HMAC. Defaults to SHA1.
		Algorithm: otp.AlgorithmSHA1,
	}
	return totp.GenerateCodeCustom(secret, time.Now(), option)
}
