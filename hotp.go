package otp

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/binary"
	"errors"
)

// GenerateHOTPCode generates otp code based on `secret` and uses counter as its moving factor.
func GenerateHOTPCode(secret []byte, counter uint64, d Digit) (int32, error) {
	counterBytes := make([]byte, 8) // 8 bytes as defined in the standard RFC-4226
	binary.BigEndian.PutUint64(counterBytes, counter)

	hmacSha1 := hmac.New(sha1.New, secret)

	_, _ = hmacSha1.Write(counterBytes)

	sum := hmacSha1.Sum(nil)

	return DynamicTruncation(sum, d)
}

// ValidateHOTPCode validates a OTP code by comparing it with the generated code
// using secret and the counter as its moving factor.
func ValidateHOTPCode(code int32, secret []byte, counter uint64, d Digit) (bool, error) {
	expected, err := GenerateHOTPCode(secret, counter, d)
	if err != nil {
		return false, err
	}

	// use constant time to prevent timing attack
	return subtle.ConstantTimeEq(expected, code) == 1, nil
}

// DynamicTruncation as defined in the rfc4226#section-5.4
func DynamicTruncation(b []byte, d Digit) (int32, error) {
	if len(b) != 20 {
		return 0, errors.New("invalid hash length")
	}

	offset := b[19] & 0xf
	binCode := (uint32(b[offset])&0x7f)<<24 |
		(uint32(b[offset+1])&0xff)<<16 |
		(uint32(b[offset+2])&0xff)<<8 |
		(uint32(b[offset+3]) & 0xff)

	return int32(binCode % d.Modulus()), nil
}
