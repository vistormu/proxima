package cmd

import (
	"github.com/vistormu/go-dsa/errors"
)

const (
	WrongExtension errors.ErrorType = "wrong file extension"
	ReadFile       errors.ErrorType = "error reading file"
	WriteFile      errors.ErrorType = "error writing file"
)
