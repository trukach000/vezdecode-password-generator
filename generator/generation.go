package generator

import (
	"math/rand"
	"strings"
)

func (gen *Generator) Generate() string {
	var pass strings.Builder

	totalSet := lowerCharSet
	if gen.IsNeedUppercase {
		totalSet = totalSet + upperCharSet
	}

	if gen.IsNeedDigits {
		totalSet = totalSet + digitSet
	}

	if gen.IsNeedSpecialCharacters {
		totalSet = totalSet + specialCharSet
	}

	for i := int64(0); i < gen.RequiredLength; i++ {
		random := rand.Intn(len(totalSet))
		pass.WriteString(string(totalSet[random]))
	}

	return pass.String()
}
