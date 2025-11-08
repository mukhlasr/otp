package otp_test

import (
	"testing"

	"github.com/mukhlasr/otp"
)

func TestDynamicTruncation(t *testing.T) {
	// example from: htps://datatracker.ietf.org/doc/html/rfc4226#section-5.4
	b := []byte{0x1f, 0x86, 0x98, 0x69, 0x0e, 0x02, 0xca, 0x16, 0x61, 0x85, 0x50, 0xef, 0x7f, 0x19, 0xda, 0x8e, 0x94, 0x5b, 0x55, 0x5a}
	res, err := otp.DynamicTruncation(b, otp.OTPDigitSix)

	if err != nil {
		t.Error("unexpected error:", err)
	}

	if res != 872921 {
		t.Error("wrong truncation result:", res)
	}
}

func TestGenerateHOTPCode(t *testing.T) {
	// test case from the https://datatracker.ietf.org/doc/html/rfc4226#page-32
	secret := []byte("12345678901234567890")
	digit := otp.OTPDigitSix
	expected := []int32{755224, 287082, 359152, 969429, 338314, 254676, 287922, 162583, 399871, 520489}
	for counter, val := range expected {
		res, err := otp.GenerateHOTPCode(secret, uint64(counter), digit)
		if err != nil {
			t.Error("unexpected error:", err)
		}

		if res != val {
			t.Errorf("expecting: %d, but got: %d", val, res)
		}
	}
}

func TestValidateHOTPCode(t *testing.T) {
	// test case from the https://datatracker.ietf.org/doc/html/rfc4226#page-32
	secret := []byte("12345678901234567890")
	digit := otp.OTPDigitSix
	expected := []int32{755224, 287082, 359152, 969429, 338314, 254676, 287922, 162583, 399871, 520489}
	for counter, val := range expected {
		valid, err := otp.ValidateHOTPCode(val, secret, uint64(counter), digit)
		if err != nil {
			t.Error("unexpected error:", err)
		}

		if !valid {
			t.Errorf("expecting a valid code")
		}
	}
}
