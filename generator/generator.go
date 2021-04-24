package generator

type Generator struct {
	RequiredLength          int64
	IsNeedUppercase         bool
	IsNeedSpecialCharacters bool
	IsNeedDigits            bool
}

func NewGenerator(requiredLength int64, needUppercase, needSpecialCharacters, needDigits bool) *Generator {
	return &Generator{
		RequiredLength:          requiredLength,
		IsNeedUppercase:         needUppercase,
		IsNeedSpecialCharacters: needSpecialCharacters,
		IsNeedDigits:            needDigits,
	}
}
