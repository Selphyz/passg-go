package password

import (
	"math/rand"
	"time"
	"unicode"
)

const (
	DefaultLength = 10
)

var (
	uppercaseRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowercaseRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	numberRunes    = []rune("0123456789")
	symbolRunes    = []rune("!@#$%^&*()-_=+[]{}|;:,.<>?")
)

type Pattern struct {
	IncludeUppercase bool
	IncludeLowercase bool
	IncludeNumbers   bool
	IncludeSymbols   bool
}

func Generate(length int, pattern Pattern) string {
	rand.Seed(time.Now().UnixNano())
	var allowedRunes []rune

	if pattern.IncludeUppercase {
		allowedRunes = append(allowedRunes, uppercaseRunes...)
	}
	if pattern.IncludeLowercase {
		allowedRunes = append(allowedRunes, lowercaseRunes...)
	}
	if pattern.IncludeNumbers {
		allowedRunes = append(allowedRunes, numberRunes...)
	}
	if pattern.IncludeSymbols {
		allowedRunes = append(allowedRunes, symbolRunes...)
	}

	if len(allowedRunes) == 0 {
		allowedRunes = append(allowedRunes, lowercaseRunes...)
		allowedRunes = append(allowedRunes, uppercaseRunes...)
		allowedRunes = append(allowedRunes, numberRunes...)
	}

	password := make([]rune, length)
	for i := range password {
		password[i] = allowedRunes[rand.Intn(len(allowedRunes))]
	}

	return string(password)
}

func IsValidPassword(s string) bool {
	var hasUpper, hasLower, hasNumber, hasSymbol bool
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}
