package otp_test

import (
	"testing"

	"github.com/mukhlasr/otp"
)

func TestZeroFill(t *testing.T) {
	type testCase struct {
		Digit  otp.Digit
		Code   int32
		Result string
	}

	expected := []testCase{
		{Digit: otp.DigitEight, Code: 1, Result: "00000001"},
		{Digit: otp.DigitEight, Code: 12, Result: "00000012"},
		{Digit: otp.DigitEight, Code: 123, Result: "00000123"},
		{Digit: otp.DigitEight, Code: 1234, Result: "00001234"},
		{Digit: otp.DigitEight, Code: 12345, Result: "00012345"},
		{Digit: otp.DigitEight, Code: 123456, Result: "00123456"},
		{Digit: otp.DigitEight, Code: 1234567, Result: "01234567"},
		{Digit: otp.DigitEight, Code: 12345678, Result: "12345678"},

		{Digit: otp.DigitSix, Code: 1, Result: "000001"},
		{Digit: otp.DigitSix, Code: 12, Result: "000012"},
		{Digit: otp.DigitSix, Code: 123, Result: "000123"},
		{Digit: otp.DigitSix, Code: 1234, Result: "001234"},
		{Digit: otp.DigitSix, Code: 12345, Result: "012345"},
		{Digit: otp.DigitSix, Code: 123456, Result: "123456"},
	}

	for _, val := range expected {
		res := otp.ZeroFill(val.Code, val.Digit)
		if res != val.Result {
			t.Errorf("expecting: %s, but got: %s", val.Result, res)
		}
	}
}
