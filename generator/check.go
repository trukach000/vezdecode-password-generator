package generator

import (
	"fmt"
	"unicode/utf8"
)

func (gen *Generator) CheckPassword(pass string) error {
	length := utf8.RuneCountInString(pass)
	if int64(length) < gen.RequiredLength {
		return NewCheckError(
			fmt.Sprintf("length of password is %d, required: %d", length, gen.RequiredLength),
		)
	}

	if gen.IsNeedUppercase {
		hasUppercase := false
		for _, r := range pass {
			if isUpper(r) {
				hasUppercase = true
				break
			}
		}
		if !hasUppercase {
			return NewCheckError("password has no uppercase")
		}
	}

	if gen.IsNeedDigits {
		hasDigit := false
		for _, r := range pass {
			if isDigit(r) {
				hasDigit = true
				break
			}
		}
		if !hasDigit {
			return NewCheckError("password has no digits")
		}
	}

	if gen.IsNeedSpecialCharacters {
		hasSpecChar := false
		for _, r := range pass {
			if isSpecialCharacter(r) {
				hasSpecChar = true
				break
			}
		}
		if !hasSpecChar {
			return NewCheckError("password has no special characters")
		}
	}

	return nil
}
