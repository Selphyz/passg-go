package tests

import (
	"github.com/Selphyz/passg/pkg/password"
	"testing"
	"unicode"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		pattern  password.Pattern
		expected func(string) bool
	}{
		{
			name:   "Default pattern",
			length: 10,
			pattern: password.Pattern{
				IncludeUppercase: true,
				IncludeLowercase: true,
				IncludeNumbers:   true,
			},
			expected: func(s string) bool {
				return len(s) == 10 &&
					containsAny(s, unicode.IsUpper) &&
					containsAny(s, unicode.IsLower) &&
					containsAny(s, unicode.IsNumber)
			},
		},
		{
			name:   "Only lowercase",
			length: 8,
			pattern: password.Pattern{
				IncludeLowercase: true,
			},
			expected: func(s string) bool {
				return len(s) == 8 && containsOnly(s, unicode.IsLower)
			},
		},
		{
			name:   "Uppercase and symbols",
			length: 12,
			pattern: password.Pattern{
				IncludeUppercase: true,
				IncludeSymbols:   true,
			},
			expected: func(s string) bool {
				return len(s) == 12 &&
					containsAny(s, unicode.IsUpper) &&
					containsAny(s, func(r rune) bool {
						return unicode.IsPunct(r) || unicode.IsSymbol(r)
					})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := password.Generate(tt.length, tt.pattern)
			if !tt.expected(result) {
				t.Errorf("Generate() = %v, doesn't match expected pattern", result)
			}
		})
	}
}

func containsAny(s string, f func(rune) bool) bool {
	for _, r := range s {
		if f(r) {
			return true
		}
	}
	return false
}

func containsOnly(s string, f func(rune) bool) bool {
	for _, r := range s {
		if !f(r) {
			return false
		}
	}
	return true
}
