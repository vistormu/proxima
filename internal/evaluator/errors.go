package evaluator

import (
	"github.com/vistormu/go-dsa/errors"
)

const (
	Script          errors.ErrorType = "error executing python component"
	MissingDef      errors.ErrorType = "missing function definition"
	NameMismatch    errors.ErrorType = "name mismatch"
	InterpreterInit errors.ErrorType = "error initializing python interpreter"
	InterpreterEval errors.ErrorType = "error evaluating python script"

	MissingComponents  errors.ErrorType = "missing component definitions"
	ReadDir            errors.ErrorType = "error reading directory"
	ReadFile           errors.ErrorType = "error reading file"
	DuplicateComponent errors.ErrorType = "duplicate component"
)
