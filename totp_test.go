package otp_test

import (
	"testing"
	"time"

	"github.com/mukhlasr/otp"
)

func TestGenerateTOTPCode(t *testing.T) {
	// test case from the https://datatracker.ietf.org/doc/html/rfc4226#page-32
	secret := []byte("12345678901234567890")
	digit := otp.OTPDigitEight
	expected := []struct {
		Time time.Time
		Code int32
	}{
		{
			Time: time.Unix(59, 0),
			Code: 94287082,
		},
		{
			Time: time.Unix(1111111109, 0),
			Code: 7081804,
		},
		{
			Time: time.Unix(1111111111, 0),
			Code: 14050471,
		},
		{
			Time: time.Unix(1234567890, 0),
			Code: 89005924,
		},
		{
			Time: time.Unix(2000000000, 0),
			Code: 69279037,
		},
		{
			Time: time.Unix(20000000000, 0),
			Code: 65353130,
		},
	}
	for _, val := range expected {
		res, err := otp.GenerateTOTPCode(secret, val.Time, digit)
		if err != nil {
			t.Error("unexpected error:", err)
		}

		if res != val.Code {
			t.Errorf("expecting: %d, but got: %d", val.Code, res)
		}
	}
}

func TestValidateTOTPCode(t *testing.T) {
	// test case from the https://datatracker.ietf.org/doc/html/rfc4226#page-32
	secret := []byte("12345678901234567890")
	digit := otp.OTPDigitEight
	expected := []struct {
		Time time.Time
		Code int32
	}{
		{
			Time: time.Unix(59, 0),
			Code: 94287082,
		},
		{
			Time: time.Unix(1111111109, 0),
			Code: 7081804,
		},
		{
			Time: time.Unix(1111111111, 0),
			Code: 14050471,
		},
		{
			Time: time.Unix(1234567890, 0),
			Code: 89005924,
		},
		{
			Time: time.Unix(2000000000, 0),
			Code: 69279037,
		},
		{
			Time: time.Unix(20000000000, 0),
			Code: 65353130,
		},
	}

	for _, val := range expected {
		valid, err := otp.ValidateTOTPCode(val.Code, secret, val.Time, digit)
		if err != nil {
			t.Error("unexpected error:", err)
		}

		if !valid {
			t.Errorf("expecting a valid code")
		}
	}
}
