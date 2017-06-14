# Go SEPA

Go SEPA provides tools and utilities for generating SEPA-XML Files and validating
the input.


## Installation

```bash
go get github.com/nirnanaaa/go-sepa
```

## Usage

Verifying Creditor Identifier:

```go
sanitized, err := sepa.SanitizeCreditorIdentifier("DE98ZZZ09999999999")
```
