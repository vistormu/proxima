package parser

import (
	"github.com/vistormu/go-dsa/errors"
)

const (
	ExpectedToken errors.ErrorType = "expected token"
	WrongTagName  errors.ErrorType = "wrong tag name"
	UnclosedTag   errors.ErrorType = "unclosed tag"
)
