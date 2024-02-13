package idam

// Password requirements
const (
	// Minimum password length
	MinPasswordLength = 8
	// Maximum password length
	MaxPasswordLength = 64
	// Allowable (require one) special characters for passwords
	AllowablePasswordSpecialCharacters = "!@#$%^&*()_+-=[]{}|\\:,./?"
	// Disallowed special characters for passwords
	DisallowedPassowrdSpecialCharacters = "\"'`~<>;"
)

// The identity provider names
const (
	MarshallLabsIdentityProvider = "marshall-labs"
)
