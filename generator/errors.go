package generator

import "fmt"

func NewCheckError(errText string) *CheckError {
	return &CheckError{ErrorText: errText}
}

type CheckError struct {
	ErrorText string
}

func (e *CheckError) Error() string {
	return fmt.Sprintf("Check failed: %s", e.ErrorText)
}
