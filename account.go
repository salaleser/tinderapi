package tinderapi

type Account struct {
	Username           string `json:"username"`
	AccountPhoneNumber string `json:"account_phone_number"`
	AppleIDLinked      bool   `json:"apple_id_linked"`
	FacebookIDLinked   bool   `json:"facebook_id_linked"`
	IsEmailVerified    bool   `json:"is_email_verified"`
	AccountEmail       string `json:"account_email"`
}
