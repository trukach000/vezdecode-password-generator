package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneration(t *testing.T) {
	tests := []struct {
		Title                   string
		RequiredLength          int64
		IsNeedUppercase         bool
		IsNeedSpecialCharacters bool
		IsNeedDigits            bool
		ExpectationLength       int
	}{
		{
			Title:             "Generate pass with zero length",
			RequiredLength:    0,
			ExpectationLength: 0,
		},
		{
			Title:             "Generate pass with negative length",
			RequiredLength:    -5,
			ExpectationLength: 0,
		},
		{
			Title:                   "Generate pass with correct length and all sets",
			RequiredLength:          10,
			ExpectationLength:       10,
			IsNeedSpecialCharacters: true,
			IsNeedDigits:            true,
			IsNeedUppercase:         true,
		},
	}

	for _, test := range tests {

		gen := NewGenerator(
			test.RequiredLength,
			test.IsNeedUppercase,
			test.IsNeedSpecialCharacters,
			test.IsNeedDigits,
		)

		passResult := gen.Generate()

		t.Log(test.Title)
		assert.Equal(t, len(passResult), test.ExpectationLength)
	}
}
