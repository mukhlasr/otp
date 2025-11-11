package otp

import (
	"time"
)

const (
	totpDefaultPeriod = 30
)

// GetTOTPCounter converts time t into a counter that can be used as the moving factor
// for the HOTP.
func GetTOTPCounter(t time.Time, period int64) uint64 {
	sec := t.Unix()
	return uint64(sec / period)
}

// GenerateTOTPCode calls GenerateTOTPCodeAlgo with algo = AlgoSHA1
func GenerateTOTPCode(secret []byte, t time.Time, d Digit) (int32, error) {
	return GenerateHOTPCode(secret, GetTOTPCounter(t, totpDefaultPeriod), d)
}

// GenerateTOTPCodeAlgo generates an otp code based on `secret` using
// HMAC-<Algo> hashing algorithm and `floor(t/30)` as its moving
// factor.
func GenerateTOTPCodeAlgo(a Algo, secret []byte, t time.Time, d Digit) (int32, error) {
	return GenerateHOTPCodeAlgo(a, secret, GetTOTPCounter(t, totpDefaultPeriod), d)
}

// ValidateTOTPCode calls ValidateTOTPCodeAlgo with algo = AlgoSHA1
func ValidateTOTPCode(passcode int32, secret []byte, t time.Time, d Digit) (bool, error) {
	return ValidateHOTPCode(passcode, secret, GetTOTPCounter(t, totpDefaultPeriod), d)
}

// ValidateTOTPCodeAlgo validates an otp code based on `secret` using
// HMAC-<Algo> hashing algorithm and `floor(t/30)` as its moving
// factor.
func ValidateTOTPCodeAlgo(a Algo, passcode int32, secret []byte, t time.Time, d Digit) (bool, error) {
	return ValidateHOTPCodeAlgo(a, passcode, secret, GetTOTPCounter(t, totpDefaultPeriod), d)
}
