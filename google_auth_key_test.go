package otp_test

import (
	"errors"
	"io"
	"testing"

	"github.com/mukhlasr/otp"
)

type dummyReader string

func (r dummyReader) Read(b []byte) (int, error) {
	n := 0

	if string(r) == "error" {
		return n, errors.New("unexpected error")
	}

	str := string(r)
	for i := range len(b) {
		b[i] = str[i%len(str)]
		n++
	}

	return n, nil
}

func TestGoogleAuthKeyString(t *testing.T) {
	t.Run("test new key", func(t *testing.T) {
		secretByteLength := 10
		testCases := []struct {
			Name         string
			RandomReader io.Reader
			Expectation  func(t *testing.T, key otp.GoogleAuthKey, err error)
		}{
			{
				Name:         "OK",
				RandomReader: dummyReader("\xDE\xEAD\xBE\xEF"),
				Expectation: func(t *testing.T, key otp.GoogleAuthKey, err error) {
					if err != nil {
						t.Errorf("expecting no err")
					}

					if len(key.Secret) != secretByteLength*8/5 {
						t.Error("wrong secret length", len(key.Secret))
					}

					if key.Secret != "33VEJPXP33VEJPXP" {
						t.Error("wrong secret:", key.Secret)
					}
				},
			},
			{
				Name:         "error",
				RandomReader: dummyReader("error"),
				Expectation: func(t *testing.T, key otp.GoogleAuthKey, err error) {
					if err == nil {
						t.Errorf("expecting error")
					}
				},
			},
		}

		for _, testcase := range testCases {
			key, err := otp.GenerateGoogleAuthKey(otp.GoogleAuthKeyParam{
				Issuer:      "Example",
				AccountName: "alice@wonderland.com",
				Type:        "totp",

				SecretByteLength: uint32(secretByteLength),
				RandReader:       testcase.RandomReader,
			})

			testcase.Expectation(t, key, err)
		}
	})

	t.Run("test encode", func(t *testing.T) {
		key := otp.GoogleAuthKey{
			GoogleAuthKeyParam: otp.GoogleAuthKeyParam{
				Issuer:           "Example",
				AccountName:      "alice@wonderland.com",
				SecretByteLength: 10,
				Type:             otp.TypeTOTP,
			},
			Secret: "SUPERSECRETKEY",
		}

		if key.String() != "otpauth://totp/Example:alice@wonderland.com?secret=SUPERSECRETKEY&issuer=Example&algorithm=SHA1&digits=6&period=30" {
			t.Error("invalid key string", key.String())
		}

		t.Run("no issuer", func(t *testing.T) {
			key := otp.GoogleAuthKey{
				GoogleAuthKeyParam: otp.GoogleAuthKeyParam{
					AccountName:      "alice@wonderland.com",
					SecretByteLength: 10,
					Type:             otp.TypeTOTP,
				},
				Secret: "SUPERSECRETKEY",
			}

			if key.String() != "otpauth://totp/:alice@wonderland.com?secret=SUPERSECRETKEY&algorithm=SHA1&digits=6&period=30" {
				t.Error("invalid key string", key.String())
			}
		})

		t.Run("no account name", func(t *testing.T) {
			key := otp.GoogleAuthKey{
				GoogleAuthKeyParam: otp.GoogleAuthKeyParam{
					Issuer:           "Example",
					SecretByteLength: 10,
					Type:             otp.TypeTOTP,
				},
				Secret: "SUPERSECRETKEY",
			}

			if key.String() != "otpauth://totp/Example:?secret=SUPERSECRETKEY&issuer=Example&algorithm=SHA1&digits=6&period=30" {
				t.Error("invalid key string", key.String())
			}
		})
	})
}
