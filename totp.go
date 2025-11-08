package otp

import (
	"math"
	"time"
)

// GetTOTPCounter converts time t into a counter that can be used as the moving factor
// for the HOTP.
func GetTOTPCounter(t time.Time) uint64 {
	sec := t.Unix()
	return uint64(math.Floor(float64(sec) / float64(period)))
}

// GenerateTOTPCode generates otp code based on `secret` and uses time `t“
// that's converted to a counter as its moving factor.
func GenerateTOTPCode(secret []byte, t time.Time, d OTPDigit) (int32, error) {
	return GenerateHOTPCode(secret, GetTOTPCounter(t), d)
}

// ValidateTOTPCode validates a OTP code by comparing it with the generated
// code using secret and time t that's converted to a counter as its moving
// factor.
func ValidateTOTPCode(passcode int32, secret []byte, t time.Time, d OTPDigit) (bool, error) {
	return ValidateHOTPCode(passcode, secret, GetTOTPCounter(t), d)
}
