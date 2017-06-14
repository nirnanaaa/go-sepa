package sepa_test

import (
	"testing"

	sepa "github.com/nirnanaaa/go-sepa"
)

func TestVersions(t *testing.T) {
	if sepa.SepaPain00100203 != 100203 {
		t.Fatal("SepaPain00100203 had wrong version: %v", sepa.SepaPain00100203)
	}
	if sepa.SepaPain00100303 != 100303 {
		t.Fatal("SepaPain00100303 had wrong version: %v", sepa.SepaPain00100303)
	}
	if sepa.SepaPain00100103 != 100103 {
		t.Fatal("SepaPain00100103 had wrong version: %v", sepa.SepaPain00100103)
	}
	if sepa.SepaPain00100103Gbic != 1001031 {
		t.Fatal("SepaPain00100103Gbic had wrong version: %v", sepa.SepaPain00100103Gbic)
	}
	if sepa.SepaPain00800202 != 800202 {
		t.Fatal("SepaPain00800202 had wrong version: %v", sepa.SepaPain00800202)
	}
	if sepa.SepaPain00800302 != 800302 {
		t.Fatal("SepaPain00800302 had wrong version: %v", sepa.SepaPain00800302)
	}
	if sepa.SepaPain00800102 != 800102 {
		t.Fatal("SepaPain00800102 had wrong version: %v", sepa.SepaPain00800102)
	}
	if sepa.SepaPain00800102Gbic != 8001021 {
		t.Fatal("SepaPain00800102Gbic had wrong version: %v", sepa.SepaPain00800102Gbic)
	}
	if sepa.SepaPain00800102Austrian003 != 8001022 {
		t.Fatal("SepaPain00800102Austrian003 had wrong version: %v", sepa.SepaPain00800102Austrian003)
	}

}
