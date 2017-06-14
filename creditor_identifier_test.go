package sepa_test

import (
	"testing"

	sepa "github.com/nirnanaaa/go-sepa"
)

const validCredIdentifier = "DE10ZZZ00000758814"
const validCredIdentifier1 = "DE98ZZZ09999999999"
const invalidCredIdentifier = "__10ZZZ00000758814"

func TestWhitespaceSanitization(t *testing.T) {
	str, err := sepa.SanitizeCreditorIdentifier(validCredIdentifier + " ")
	if err != nil {
		t.Fatal(err.Error())
	}
	if str != validCredIdentifier {
		t.Fatal("did not remote whitespaces")
	}
}
func TestInvalidCreditorIdentifier(t *testing.T) {
	_, err := sepa.SanitizeCreditorIdentifier(invalidCredIdentifier)
	if err == nil {
		t.Fatal("regex on creditor identifier didn't match.")
	}
}
func TestCapitalization(t *testing.T) {
	str, err := sepa.SanitizeCreditorIdentifier("de10ZZZ00000758814")
	if err != nil {
		t.Fatal(err.Error())
	}
	if str != validCredIdentifier {
		t.Fatal("did not remote whitespaces")
	}
}

func TestStringToIntConversion(t *testing.T) {
	str := sepa.ConvertStringToMatchingInts(validCredIdentifier)
	comp := "13141035353500000758814"
	if str != comp {
		t.Fatal("alphabet mappings were done poorly.")
	}
}

func TestChecksums(t *testing.T) {
	comp := "00000758814131410"
	str := sepa.Iso7064Mod97m10ChecksumCheck(comp)
	if !str {
		t.Fatalf("expected true, got %v", str)
	}
}

func TestBadFormatting(t *testing.T) {
	comp := "d e98 Z ZZ 09 99 9999999"
	str, err := sepa.SanitizeCreditorIdentifier(comp)
	if err != nil {
		t.Fatal(err.Error())
	}
	if str != validCredIdentifier1 {
		t.Fatal("Creditor ID wasn't sanitized properly")
	}
}

func TestInvalidChars(t *testing.T) {
	comp := "DE98ZZÄ09999999998"
	_, err := sepa.SanitizeCreditorIdentifier(comp)
	if err == nil {
		t.Fatal("regex on creditor identifier didn't match.")
	}
}

func TestTooShortInput(t *testing.T) {
	comp := "DE9"
	_, err := sepa.SanitizeCreditorIdentifier(comp)
	if err == nil {
		t.Fatal("program didn't accept garbage.")
	}
}

func TestInvalidChecksum(t *testing.T) {
	comp := "DE98ZZZ09999999998"
	_, err := sepa.SanitizeCreditorIdentifier(comp)
	if err == nil {
		t.Fatal("invalid checksum was allowed.")
	}
}

func TestIgnoreMiddlePart(t *testing.T) {
	comp := "DE98abc09999999999"
	str, err := sepa.SanitizeCreditorIdentifier(comp)
	if err != nil {
		t.Fatal(err.Error())
	}
	if str != "DE98ABC09999999999" {
		t.Fatal("Creditor ID wasn't sanitized properly")
	}
}
