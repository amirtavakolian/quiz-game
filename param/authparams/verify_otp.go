package authparams

type VerifyParam struct {
	PhoneNumber string `json:"phone_number"`
	OTPCode     string `json:"otpcode"`
}
