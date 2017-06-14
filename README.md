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
sanitized, err := sepa.SanitizeCreditorIdentifier("DE98ZZZ09999999999")
```
