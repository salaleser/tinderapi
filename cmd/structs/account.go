package structs

type Account struct {
	Username           string `json:"username,omitempty"`
	AccountPhoneNumber string `json:"account_phone_number,omitempty"`
	AppleIDLinked      bool   `json:"apple_id_linked,omitempty"`
	FacebookIDLinked   bool   `json:"facebook_id_linked,omitempty"`
	IsEmailVerified    bool   `json:"is_email_verified,omitempty"`
	AccountEmail       string `json:"account_email,omitempty"`
}
