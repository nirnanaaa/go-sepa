package sepa_test

import (
	"testing"

	sepa "github.com/nirnanaaa/go-sepa"
)

func TestCorrectBicEmptyAllowed(t *testing.T) {
	id := sepa.NewBic("")
	_, err := id.Validate(true)
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestCorrectBicEmptyNowAllowed(t *testing.T) {
	id := sepa.NewBic("")
	_, err := id.Validate(false)
	if err == nil {
		t.Fatal("Empty bic should not have been allowed")
	}
}
func TestCorrectBicShort(t *testing.T) {
	comp := "COBADEFF"
	id := sepa.NewBic(comp)
	str, err := id.Validate(true)
	if err != nil {
		t.Fatal(err.Error())
	}
	if str != "COBADEFFXXX" {
		t.Fatal("BIC wasn't sanitized properly")
	}
}

func TestCorrectBicLong(t *testing.T) {
	comp := "COBADEFFXXX"
	id := sepa.NewBic(comp)
	str, err := id.Validate(true)
	if err != nil {
		t.Fatal(err.Error())
	}
	if str != "COBADEFFXXX" {
		t.Fatal("BIC wasn't sanitized properly")
	}
}

func TestCorrectBicLongDowncased(t *testing.T) {
	comp := "cobadeffxxx"
	id := sepa.NewBic(comp)
	str, err := id.Validate(true)
	if err != nil {
		t.Fatal(err.Error())
	}
	if str != "COBADEFFXXX" {
		t.Fatal("BIC wasn't sanitized properly")
	}
}

func TestCorrectBicStringConversion(t *testing.T) {
	comp := "cobadeffxxx"
	id := sepa.NewBic(comp)
	_, err := id.Validate(true)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(id.String())
	if id.String() != "COBADEFFXXX" {
		t.Fatal("BIC wasn't sanitized properly")
	}
}
