package sepa

import (
	"fmt"
	"regexp"
)

const errMesgIBANInvalid = ""

// IBAN defines a struct to work on IBANs
type IBAN struct {
	CountryCode   string
	Checksum      string
	AccountNumber string
	BBAN          string
}

// NewIBAN creates a new IBAN object
func NewIBAN(iban string) (*IBAN, error) {
	sanitized := UniformString(iban)
	valid := validateIBANRegex(sanitized)
	if !valid {
		return nil, fmt.Errorf(errMesgIBANInvalid)
	}
	cc := CountryCode(sanitized)
	checksum := getIbanChecksum(sanitized)
	accountNumber := getAccountNumberIBAN(sanitized)
	bankID := getBankIDIban(sanitized)
	return &IBAN{cc, checksum, accountNumber, bankID}, nil
}

// Validate checks if an IBAN is valid. Default check is based on checksum.
func (i *IBAN) Validate() (string, error) {
	return "", nil
}

// ValidateRegex checks if an IBAN is valid based on its checksum
func (i *IBAN) ValidateRegex() (string, error) {
	return "", nil
}

func (i *IBAN) String() string {
	return fmt.Sprintf("%s%s%s%s", i.CountryCode, i.Checksum, i.BBAN, i.AccountNumber)
}

// IsNationalTo validates if the provided iban is in the same country as the one constructed
func (i *IBAN) IsNationalTo(iban string) bool {
	return i.CountryCode == CountryCode(iban)
}

func validateIBANRegex(input string) bool {
	re := regexp.MustCompile(PatternIban)
	return re.MatchString(input)
}
func getAccountNumberIBAN(str string) string {
	return str[13:]
}
func getBankIDIban(str string) string {
	return str[4:13]
}
func getIbanChecksum(str string) string {
	return str[2:4]
}

// CountryCode returns the country code for a sepa formatted string
func CountryCode(str string) string {
	return str[:2]
}
