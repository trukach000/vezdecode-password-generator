package generator

const lowerCharSet = "abcdefghijklmnopqrstuvwxyz"
const upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const specialCharSet = "!@#$%&*_"
const digitSet = "0123456789"

func isUpper(e int32) bool {
	for _, a := range upperCharSet {
		if a == e {
			return true
		}
	}
	return false
}

func isSpecialCharacter(e int32) bool {
	for _, a := range specialCharSet {
		if a == e {
			return true
		}
	}
	return false
}

func isDigit(e int32) bool {
	for _, a := range digitSet {
		if a == e {
			return true
		}
	}
	return false
}
