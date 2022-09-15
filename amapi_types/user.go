package amapi_types

import "time"

type User struct {
	Meta          *Meta         `json:"meta"`
	Eppn          string        `json:"eppn"`    //`json:"eduPersonPrincipalName"`
	UserID        string        `json:"user_id"` // bson.ObjectID
	GivenName     string        `json:"given_name"`
	DisplayName   string        `json:"display_name"`
	Surname       string        `json:"surname"`
	Subject       string        `json:"subject"`
	Language      string        `json:"language"`
	MailAddresses MailAddresses `json:"mail_addresses"`
	PhoneNumbers  PhoneNumbers  `json:"phone_numbers"`
	Credentials   Credentials   `json:"credentials"`
	Identities    Identities    `json:"identities"`
	ModifiedTS    time.Time     `json:"modified_ts"`
	Entitlements  []string      `json:"entitlements"`
	//	tou
	Terminated time.Time
	//	locked_identity
	//	orcid
	Ladok *Ladok `json:"ladok"`
	//	profiles
	//	letter_proofing_data
	RevokedTS time.Time
}

type Meta struct {
	Cleaned    map[string]time.Time `json:"cleaned"`
	ModifiedTS time.Time            `json:"modified_ts"`
	Version    string               `json:"version"` // bson.ObjectID
}

type MailAddresses []*MailAddress

type MailAddress struct {
	CreatedBy  string    `json:"created_by"`
	CreatedTS  time.Time `json:"created_ts"`
	Email      string    `json:"email"`
	Primary    bool      `json:"primary"`
	Verified   bool      `json:"verified"`
	VerifiedBy string    `json:"verified_by"`
	VerifiedTS time.Time `json:"verified_ts"`
}

type PhoneNumbers []*PhoneNumber

type PhoneNumber struct {
	CreatedBy  string    `json:"created_by"`
	CreatedTS  time.Time `json:"created_ts"`
	Number     string    `json:"number"`
	Primary    bool      `json:"primary"`
	Verified   bool      `json:"verified"`
	VerifiedBy string    `json:"verified_by"`
	VerifiedTS time.Time `json:"verified_ts"`
}

type Identities []*Identity

type Identity struct {
	CreatedBy    string    `json:"created_by"`
	CreatedTS    time.Time `json:"created_ts"`
	IdentityType string    `json:"identity_type"`
	ModifiedTS   time.Time `json:"modified_ts"`
	Number       string    `json:"number"`
	Verified     bool      `json:"verified"`
	VerifiedBy   string    `json:"verified_by"`
	VerifiedTS   time.Time `json:"verified_ts"`
}

type Credentials []*Credential

type Credential struct {
	CreatedBy    string    `json:"created_by"`
	CreatedTS    time.Time `json:"created_ts"`
	CredentialID string    `json:"credential_id"`
	IsGenerated  bool      `json:"is_generated"`
	ModifiedTS   time.Time `json:"modified_ts"`
	Salt         string    `json:"salt"`
}

type UniversityName struct {
	EN string `json:"en"`
	SV string `json:"sv"`
}

type University struct {
	CreatedTS  time.Time      `json:"created_ts"`
	LadokName  string         `json:"ladok_name"`
	ModifiedTS time.Time      `json:"modified_ts"`
	Name       UniversityName `json:"name"`
	Verified   bool           `json:"verified"`
	VerifiedBy string         `json:"verified_by"`
}

type Ladok struct {
	CreatedTS  time.Time  `json:"created_ts"`
	ExternalID string     `json:"external_id"`
	ModifiedTS time.Time  `json:"modified_ts"`
	University University `json:"university"`
}
