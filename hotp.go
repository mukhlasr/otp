package otp

import (
	"crypto/hmac"
	"crypto/subtle"
	"encoding/binary"
	"errors"
)

// GenerateHOTPCode calls GenerateHOTPCodeAlgo with algo = AlgoSHA1
func GenerateHOTPCode(secret []byte, counter uint64, d Digit) (int32, error) {
	return GenerateHOTPCodeAlgo(AlgoSHA1, secret, counter, d)
}

// GenerateHOTPCodeAlgo generates an otp code based on `secret` using
// HMAC-<Algo> hashing algorithm and `counter` as its moving factor.
func GenerateHOTPCodeAlgo(a Algo, secret []byte, counter uint64, d Digit) (int32, error) {
	counterBytes := make([]byte, 8) // 8 bytes as defined in the standard RFC-4226
	binary.BigEndian.PutUint64(counterBytes, counter)

	hmacSha1 := hmac.New(a.Hash, secret)

	_, _ = hmacSha1.Write(counterBytes)

	sum := hmacSha1.Sum(nil)

	return DynamicTruncation(sum, d)
}

// ValidateHOTPCode calls ValidateHOTPCodeAlgo with algo = AlgoSHA1
func ValidateHOTPCode(code int32, secret []byte, counter uint64, d Digit) (bool, error) {
	return ValidateHOTPCodeAlgo(AlgoSHA1, code, secret, counter, d)
}

// ValidateHOTPCodeAlgo validates an otp code based on `secret` using
// HMAC-<Algo> hashing algorithm and `counter` as its moving factor.
func ValidateHOTPCodeAlgo(a Algo, code int32, secret []byte, counter uint64, d Digit) (bool, error) {
	expected, err := GenerateHOTPCodeAlgo(a, secret, counter, d)
	if err != nil {
		return false, err
	}

	// use constant time to prevent timing attack
	return subtle.ConstantTimeEq(expected, code) == 1, nil
}

// DynamicTruncation as defined in the rfc4226#section-5.4
func DynamicTruncation(b []byte, d Digit) (int32, error) {
	if len(b) < 20 {
		return 0, errors.New("invalid byte length")
	}

	offset := b[len(b)-1] & 0xf
	binCode := (uint32(b[offset])&0x7f)<<24 |
		(uint32(b[offset+1])&0xff)<<16 |
		(uint32(b[offset+2])&0xff)<<8 |
		(uint32(b[offset+3]) & 0xff)

	return int32(binCode % d.Modulus()), nil
}
