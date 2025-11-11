package otp

import (
	"fmt"
	"io"
	"net/url"
)

const (
	googleAuthDigit = DigitSix
)

// GoogleAuthKeyParam defines the structure of otp key supported by Google
// Authenticator App.
type GoogleAuthKeyParam struct {
	// Issuer of the otp
	Issuer string
	// AccountName of the user
	AccountName string
	// Type of the OTP(hotp/totp)
	Type Type

	// SecretByteLength defines the length of bytes will be read from RandReader
	SecretByteLength uint32
	// RandReader is the reader to read random bytes from. If nil, then it will
	// use the crypto/rand.Reader.
	RandReader io.Reader
}

// GoogleAuthKey provides key that works with Google Authenticator or any
// clients that have a similar implementation.
type GoogleAuthKey struct {
	GoogleAuthKeyParam
	Secret string
}

// GenerateGoogleAuthKey generates a new usable key for Google Authenticator.
func GenerateGoogleAuthKey(p GoogleAuthKeyParam) (GoogleAuthKey, error) {
	var key GoogleAuthKey

	secret, err := GenerateBase32Secret(p.SecretByteLength, p.RandReader)
	if err != nil {
		return key, fmt.Errorf("failed to generate key: %w", err)
	}

	key.GoogleAuthKeyParam = p
	key.Secret = secret

	return key, nil
}

// String encodes the key into a string url that can be shown as a QR code and
// then scanned by Google Authenticator.
func (k *GoogleAuthKey) String() string {
	q := ""
	q = k.appendQueryIfNotEmpty(q, "secret", k.Secret)
	q = k.appendQueryIfNotEmpty(q, "issuer", k.Issuer)
	q = k.appendQueryIfNotEmpty(q, "algorithm", "SHA1")
	q = k.appendQueryIfNotEmpty(q, "digits", fmt.Sprint(googleAuthDigit))
	q = k.appendQueryIfNotEmpty(q, "period", fmt.Sprint(totpDefaultPeriod))

	u := &url.URL{
		Scheme:   "otpauth",
		Host:     string(k.Type),
		Path:     "/" + k.Issuer + ":" + k.AccountName,
		RawQuery: q,
	}

	return u.String()
}

func (k *GoogleAuthKey) appendQueryIfNotEmpty(raw, key, value string) string {
	if value == "" {
		return raw
	}

	value = url.PathEscape(value)
	q := fmt.Sprintf("%s=%s", key, value)
	if raw == "" {
		return q
	}

	return raw + "&" + q
}
