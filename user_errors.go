package api

const (
	RateLimit = "3"

	// Missing fields
	TxTypeId        = "ERR4000"
	TxReference     = "ERR4001"
	TxAmountInvalid = "ERR4002"
	TxBetId         = "ERR4003"
	TxAuthCode      = "ERR4004"
	TxIdInvalid     = "ERR4005"
	PartnerKey      = "ERR4006"
	Hash            = "ERR4007"
	TxTypeIdInvalid = "ERR4008"
	TxUserId        = "ERR4009"

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
	HashInvalid        = "ERR2028"

	// Reset tokens
	ResetTokenInvalid = "ERR2200"

	// Complexity
	PasswordComplexity = "ERR2250"
	UsernameComplexity = "ERR2260"

	// KYC
	AgeNotMet = "ERR2002"

	// User verification
	UserVerificationFail = "ERR2100"
	AddressWarning       = "ERR2101"
	UserWarning          = "ERR2102"
	UserKycFail          = "ERR2103"

	// System error
	SystemException  = "SYS1000"
	TokenInvalid     = "SYS2000"
	LoginTimeOut     = "SYS3000"
	GenericErrorCode = "SYS9999"

	// Suspension / Exclusion
	UserExcluded  = "ERR3000"
	UserSuspended = "ERR3050"

	// Not Found
	EmailNotFound = "ERR3100"

	// Failure
	TooManyLoginFails = "ERR2400"

	// Gamstop
	GamstopRegistered = "ERR2450"

	// Transactions API
	InsufficientFunds     = "ERR4050"
	TransactionNotFound   = "ERR4060"
	TransactionNotPending = "ERR4070"
)

func (s *Error) GetErrorMessage(errorCode string) string {

	errors := map[string]string{
		constants.Email:            "Missing email address",
		constants.Password:         "Missing password",
		constants.FirstName:        "Missing firstname",
		constants.LastName:         "Missing last name",
		constants.DOB:              "Missing Date of birth",
		constants.Line1:            "Missing first line of address",
		constants.City:             "Missing city",
		constants.Postcode:         "Missing postcode",
		constants.CountryIso:       "Missing country ISO Code",
		constants.PhoneNumber:      "Missing phone number",
		constants.SecurityAnswer:   "Missing security answer",
		constants.SecurityQuestion: "Missing security question",
		constants.UsernameEmail:    "Missing username / email",
		constants.TxReference:      "Missing transaction reference",
		constants.TxTypeId:         "Missing transaction type ID",
		constants.TxBetId:          "Missing bet ID",

		// Duplicates
		constants.DuplicateEmail:    "Email already in use",
		constants.DuplicateUsername: "Username already in use",

		// Invalid formats
		constants.PhoneNumberInvalid: "Invalid phone number",
		constants.EmailInvalid:       "Invalid email",
		constants.UsernameInvalid:    "Invalid username",
		constants.DOBFormat:          "Invalid date format for Date of Birth. Accepted format is yyyy-mm-dd",
		constants.DOBInvalid:         "Invalid Date of birth",
		constants.InvalidRequest:     "Invalid request",
		constants.CountryIsoInvalid:  "Invalid country ISO",
		constants.CountryNotFound:    "Specified country not found",
		constants.GenderInvalid:      "Invalid gender",
		constants.FirstnameInvalid:   "Invalid characters in firstname",
		constants.LastnameInvalid:    "Invalid characters in lastname",
		constants.LoginInvalid:       "Invalid login credentials provided",
		constants.TxAmountInvalid:    "Invalid amount",
		constants.TxIdInvalid:        "Invalid id",
		constants.TxTypeIdInvalid:    "Transaction type not supported for this request.",
		constants.HashInvalid:        "Incorrect hash provided for the given request.",
		constants.TxUserId:           "Missing user ID",

		// Password reset
		constants.ResetTokenInvalid: "Invalid or missing reset token parameter",

		// Age / DoB
		constants.AgeNotMet: "Minimum age requirement not met",

		// User verification
		constants.UserVerificationFail: "A user already exists with the details provided. If you have forgotten your password please use the password reset functionality.",
		constants.AddressWarning:       "There might already be a user registered on the same address.",
		constants.UserWarning:          "A user with the same name already exists.",
		constants.UserKycFail:          "This account has not passed KYC.",

		// System error
		constants.SystemException:  "A system error occurred. Please contact an administrator.",
		constants.TokenInvalid:     "Invalid token",
		constants.LoginTimeOut:     "Your log-in Attempt Failed, You have exceeded the maximum number of log-in attempts permitted during this session.",
		constants.GenericErrorCode: "Missing or invalid field specified",

		// Complexity
		constants.PasswordComplexity: "The password must be 8 to 100 characters long and must contain at least : one uppercase letter, one lowercase letter, one number and a special character.",
		constants.UsernameComplexity: "The username must be 4 to 100 characters long and may contain numbers, letters and special characters.",

		// Exclusion & suspension
		constants.UserExcluded:  "This account has been excluded.",
		constants.UserSuspended: "This account has been suspended.",

		// Not Found
		constants.EmailNotFound: "Email specified does not exist",

		// Too many tries
		constants.TooManyLoginFails: "Too many failed login attempts",

		// Gamstop
		constants.GamstopRegistered: "User is registered on gamstop",

		// Funds
		constants.InsufficientFunds:     "Insufficient funds",
		constants.TransactionNotFound:   "No transaction found by that Id",
		constants.TransactionNotPending: "This transaction cannot be charged because it is not pending and has already been processed.",
	}

	return errors[errorCode]
}
