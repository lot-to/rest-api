package api

const (

	// Missing fields
	Email            = "ERR1000"
	Password         = "ERR1001"
	FirstName        = "ERR1002"
	LastName         = "ERR1003"
	DOB              = "ERR1004"
	Line1            = "ERR1005"
	City             = "ERR1006"
	Postcode         = "ERR1007"
	CountryIso       = "ERR1008"
	PhoneNumber      = "ERR1009"
	SecurityAnswer   = "ERR1010"
	SecurityQuestion = "ERR1011"
	UsernameEmail    = "ERR1012"

	// Validation failures
	PhoneNumberInvalid = "ERR2040"
	EmailInvalid       = "ERR2050"
	UsernameInvalid    = "ERR2060"
	DOBFormat          = "ERR2070"
	DOBInvalid         = "ERR2080"
	InvalidRequest     = "ERR2020"
	DuplicateEmail     = "ERR2000"
	DuplicateUsername  = "ERR2001"
	CountryIsoInvalid  = "ERR2021"
	CountryNotFound    = "ERR2022"
	GenderInvalid      = "ERR2024"
	FirstnameInvalid   = "ERR2025"
	LastnameInvalid    = "ERR2026"
	LoginInvalid       = "ERR2027"

	// Reset tokens
	ResetTokenInvalid = "ERR2200"

	// Password complexity
	PasswordComplexity = "ERR2250"

	// KYC
	AgeNotMet = "ERR2002"

	// User verification
	UserVerificationFail = "ERR2100"
	AddressWarning       = "ERR2101"
	UserWarning          = "ERR2102"

	// Rate limit failure
	TooManyLoginFails = "ERR2400"

	// Suspension / Exclusion
	UserExcluded  = "ERR3000"
	UserSuspended = "ERR3050"

	// Not Found
	EmailNotFound = "ERR3100"

	// System error
	SystemException  = "SYS1000"
	TokenInvalid     = "SYS2000"
	LoginTimeOut     = "SYS3000"
	GenericErrorCode = "SYS9999"
)