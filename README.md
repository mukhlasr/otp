# OTP(HOTP/TOTP) package for Golang
This is a simple OTP(HOTP/TOTP) implementation in Go that can be used with the
Google Authenticator app. Currently, It only support HMAC-SHA1 for the hashing 
algorithm.

## Simple TOTP Example
```go
digit := otp.DigitSix
otpCode, err := otp.GenerateTOTPCode(rawByteSecret, time.Now(), digit)
fmt.Println(otp.ZeroFill(otpCode, digit))
```

## TOTP for Google Authenticator
### Generating Key
```go
// This will generate key for 6 digits OTP that can be used by Google 
// Authenticator App.
key, err := otp.GenerateGoogleAuthKey(otp.GoogleAuthKeyParam{
			Issuer:           "Example",
			AccountName:      "alice@example.com",
			Type:             otp.TypeTOTP,
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
rawByteSecret, err := otp.DecodeBase32Secret(secret)
var otpInput string
fmt.Scanf("%s", &otpInput)
code, err := strconv.Atoi(otpInput)
ok, err := otp.ValidateTOTPCode(int32(code), rawByteSecret, time.Now(), otp.DigitSix)
if ok {
	// Valid OTP
}
```