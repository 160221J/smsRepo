package apperrors

import "errors"

var ErrNotFound = errors.New("not found")
var ErrConflict = errors.New("conflict")

type ValidationError struct{ Field, Msg string }

func (e ValidationError) Error() string { return e.Field + ": " + e.Msg }
