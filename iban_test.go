package sepa_test

import (
	"testing"

	sepa "github.com/nirnanaaa/go-sepa"
)

func TestValidIBAN(t *testing.T) {
	iban := "DE21700519950000007229"
	ban, err := sepa.NewIBAN(iban)
	if err != nil {
		t.Fatal(err.Error())
	}
	if ban.String() != iban {
		t.Fatalf("Formatting iban was done poorly: got %s, expected %s", ban.String(), iban)
	}
}

func TestValidIBANFormat(t *testing.T) {
	iban := "BLAAAAAAAAAAA"
	_, err := sepa.NewIBAN(iban)
	if err == nil {
		t.Fatal("Iban is invalid but regex allowed it!")
	}
}

func TestValidIBANIsNational(t *testing.T) {
	iban := "DE21700519950000007229"
	iban1, err := sepa.NewIBAN(iban)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !iban1.IsNationalTo("DE1234") {
		t.Fatal("should be a national iban")
	}
}
