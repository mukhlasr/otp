// Package otp provides functionalities for generating and validating HOTP(IETF RFC4226)/TOTP(IETF RFC6238) codes.
package otp

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

const (
	period = 30
)

type OTPType string

const (
	OTPTypeHOTP OTPType = "hotp"
	OTPTypeTOTP OTPType = "totp"
)

type OTPDigit int

const (
	OTPDigitSix   OTPDigit = 6
	OTPDigitEight OTPDigit = 8
)

func (d OTPDigit) Modulus() uint32 {
	switch d {
	case OTPDigitSix:
		return 1_000_000
	case OTPDigitEight:
		return 100_000_000
	}

	return 0
}

// GenerateRandomSecret generate random bytes with length `byteLen` from the randomReader.
// If the randomReader is nil, it will use the default crypto/rand.Reader to read the bytes from.
func GenerateRandomSecret(byteLen uint32, randomReader io.Reader) ([]byte, error) {
	if randomReader == nil {
		randomReader = rand.Reader
	}

	b := make([]byte, byteLen)
	n, err := randomReader.Read(b)

	if err != nil {
		return b, fmt.Errorf("otp: failed to generate secret: %w", err)
	}

	if n != int(byteLen) {
		return b, errors.New("invalid random bytes length")
	}

	return b, nil

}

// GenerateBase32Secret generates an otp secret with length = `byteLen` encoded
// in base32. This is required for the otp to work with Google Authenticator.
// If the randomReader is nil, it will use the default crypto/rand.Reader to
// read the bytes from.
func GenerateBase32Secret(byteLen uint32, randomReader io.Reader) (string, error) {
	b, err := GenerateRandomSecret(byteLen, randomReader)
	return base32encoder.EncodeToString(b), err
}

// DecodeBase32Secret decodes the encoded secret and returns its raw byte form to be used by the OTP functions
func DecodeBase32Secret(secret string) ([]byte, error) {
	return base32encoder.DecodeString(secret)
}
