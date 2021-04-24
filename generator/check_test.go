package generator

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestChecking(t *testing.T) {
	tests := []struct {
		Title                   string
		Pass                    string
		RequiredLength          int64
		IsNeedUppercase         bool
		IsNeedSpecialCharacters bool
		IsNeedDigits            bool
		ExpectationCheck        bool
		ExpectationError        error
	}{
		{
			Title:            "Pass with correct length",
			Pass:             "super_pass",
			RequiredLength:   10,
			ExpectationError: nil,
		},
		{
			Title:            "Short password",
			Pass:             "123",
			RequiredLength:   10,
			ExpectationError: &CheckError{ErrorText: "length of password is 3, required: 10"},
		},
		{
			Title:            "Pass with uppercase",
			Pass:             "abCD",
			RequiredLength:   2,
			IsNeedUppercase:  true,
			ExpectationError: nil,
		},
		{
			Title:            "Pass without uppercase",
			Pass:             "abcde",
			RequiredLength:   2,
			IsNeedUppercase:  true,
			ExpectationError: &CheckError{ErrorText: "password has no uppercase"},
		},
		{
			Title:            "Pass with digits",
			Pass:             "abCD1aaa",
			RequiredLength:   2,
			IsNeedDigits:     true,
			ExpectationError: nil,
		},
		{
			Title:            "Pass without digits",
			Pass:             "abcde___",
			RequiredLength:   2,
			IsNeedDigits:     true,
			ExpectationError: &CheckError{ErrorText: "password has no digits"},
		},
		{
			Title:                   "Pass with special characters",
			Pass:                    "ab_234#%g1!",
			RequiredLength:          2,
			IsNeedSpecialCharacters: true,
			ExpectationError:        nil,
		},
		{
			Title:                   "Pass without special characters",
			Pass:                    "qwerty",
			RequiredLength:          2,
			IsNeedSpecialCharacters: true,
			ExpectationError:        &CheckError{ErrorText: "password has no special characters"},
		},
	}

	for _, test := range tests {

		gen := NewGenerator(
			test.RequiredLength,
			test.IsNeedUppercase,
			test.IsNeedSpecialCharacters,
			test.IsNeedDigits,
		)

		checkError := gen.CheckPassword(test.Pass)

		t.Log(test.Title)
		assert.Equal(t, checkError, test.ExpectationError)
	}
}
