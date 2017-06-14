package sepa_test

import (
	"testing"

	sepa "github.com/nirnanaaa/go-sepa"
)

func TestUniform(t *testing.T) {
	output := sepa.UniformString("de 12 34")
	if output != "DE1234" {
		t.Fatal("didn't format string properly")
	}
}
