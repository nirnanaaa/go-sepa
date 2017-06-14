# Go SEPA

[![Build Status](https://travis-ci.org/nirnanaaa/go-sepa.svg?branch=master)](https://travis-ci.org/nirnanaaa/go-sepa)

#

Go SEPA provides tools and utilities for generating SEPA-XML Files and validating
the input.


## Installation

```bash
go get github.com/nirnanaaa/go-sepa
```

## Usage

Verifying **creditor identifier (ci)**:

```go
ci := sepa.NewCreditorIdentifier("DE98ZZZ09999999999")
sanitized, err := ci.Validate()
```

Verifying **IBAN (iban)**:

```go
iban := sepa.NewIBAN("DE21700519950000007229")
sanitized, err := iban.Validate()
```
