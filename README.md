# OTP(HOTP/TOTP) package for Golang
This is a simple OTP(HOTP/TOTP) implementation in Go and provide support for 
Google Authenticator app. Currently, It only provide HMAC-SHA1 for the hashing 
algorithm.

## Simple TOTP Example
```go
digit := otp.OTPDigitSix
otpCode, err := otp.GenerateTOTPCode(keyBytes, time.Now(), digit)
fmt.Println(ZeroFill(otpCode, digit))
```

## TOTP for Google Authenticator
### Generating Key
```go
// This will generate key for 6 digits OTP that can be used by Google 
// Authenticator App.
key, err := otp.GenerateGoogleAuthKey(otp.GoogleAuthKeyParam{
			Issuer:           "Example",
			AccountName:      "alice@example.com",
			Type:             otp.OTPTypeTOTP,
			SecretByteLength: 10,
})

// prints the key
fmt.Println(key.String()) // it returns a url encoded otp key
// Or, generates a QR code that can be scanned by the app.
GenerateQRCode(key.String()) 
```

### Validating
```go
secret := key.Secret // get the secret from the key
digit := otp.GoogleAuthDigit
rawByteSecret, err := otp.DecodeBase32Secret(secret)
ok, err := otp.ValidateTOTPCode(code, rawByteSecret, time.Now, otp.OTPDigitSix)
if ok {
    // valid
}
```