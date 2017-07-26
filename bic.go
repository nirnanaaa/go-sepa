package sepa

import (
	"fmt"
	"regexp"
)

// BIC defines a struct to perform validation
type BIC struct {
	BIC string
}

// NewBic creates a new BIC struct
func NewBic(input string) BIC {
	return BIC{input}
}

// Validate checks if a bic is valid.
// Note that also if the bic is valid it does not have to exist
func (c BIC) Validate(allowEmpty bool) (string, error) {
	sanitized := c.String()
	if sanitized == "" && allowEmpty {
		return sanitized, nil
	}
	if c.validateWithRegex(sanitized) {
		return sanitized, nil
	}
	return "", fmt.Errorf("Input %s isn't a valid bic", sanitized)
}

func (c BIC) validateWithRegex(input string) bool {
	re := regexp.MustCompile(PatternBic)
	return re.MatchString(input)
}

func (c BIC) String() string {
	sanitized := UniformString(c.BIC)
	if len(sanitized) == 8 {
		sanitized = sanitized + "XXX"
	}
	return sanitized
}
