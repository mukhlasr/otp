# OTP(HOTP/TOTP) package for Golang
This is a simple OTP(HOTP/TOTP) implementation in Go that can be used with the
Google Authenticator app. It only support HMAC-SHA1, HMAC-SHA256, and 
HMAC-SHA512.

**Note**: All the functions to generate and validate the OTP expect the secret 
to be unencoded(raw bytes). If you want to use base32encoded secret make sure 
to decode it first by calling `otp.DecodeBase32Secret` utility function. Also, 
there is a function to generate a random base32 encoded secret 
`GenerateBase32Secret`.

## Simple OTP Example
### Basic 6 digits HMAC-SHA1 HOTP
```go
digit := otp.DigitSix
counter := uint64(0}
// generate HMAC-SHA1 TOTP code with 6 digit
otpCode, err := otp.GenerateHOTPCode(rawByteSecret, counter, digit)
fmt.Println(otp.ZeroFill(otpCode, digit))
```

### Basic 6 digits HMAC-SHA1 TOTP with 30 seconds duration
```go
digit := otp.DigitSix
// generate HMAC-SHA1 TOTP code with 6 digit and 30 seconds duration.
otpCode, err := otp.GenerateTOTPCode(rawByteSecret, time.Now(), digit)
fmt.Println(otp.ZeroFill(otpCode, digit))
```

### Custom 8 Digits OTP with HMAC-SHA256/HMAC-SHA512
```go
digit := otp.DigitEight
counter := uint64(0)
// HOTP HMAC-SHA256
otpCode, err := otp.GenerateHOTPCodeAlgo(otp.AlgoSHA256, rawByteSecret, counter, digit)
fmt.Println(otp.ZeroFill(otpCode, digit))

// TOTP HMAC-SHA512
otpCode, err = otp.GenerateTOTPCodeAlgo(otp.AlgoSHA512, rawByteSecret, time.Now(), digit)
fmt.Println(otp.ZeroFill(otpCode, digit))
```

## TOTP for Google Authenticator
### Generating Key
```go
// This will generate key for 6 digits OTP that can be used by Google 
// Authenticator App. The period duration will always be 30 seconds.
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

## TOTP with custom duration
### Generating the code
```go
counter := otp.GetTOTPCounter(time.Now(), 20)
otpCode, err := otp.GenerateHOTPCode(rawByteSecret, counter, digit)
fmt.Println(otp.ZeroFill(otpCode, digit))
```

### Validating the code
```go
var otpInput string
fmt.Scanf("%s", &otpInput)
code, err := strconv.Atoi(otpInput)
counter := otp.GetTOTPCounter(time.Now(), 20)
otpCode, err := otp.ValidateHOTPCode(rawByteSecret, counter, digit)
fmt.Println(otp.ZeroFill(otpCode, digit))
```