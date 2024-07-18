package tests

import (
	"github.com/Selphyz/passg/cmd"
	"github.com/Selphyz/passg/pkg/password"
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		expected password.Pattern
	}{
		{
			name:    "Empty pattern",
			pattern: "",
			expected: password.Pattern{
				IncludeUppercase: true,
				IncludeLowercase: true,
				IncludeNumbers:   true,
			},
		},
		{
			name:    "Alphanumeric",
			pattern: "A",
			expected: password.Pattern{
				IncludeUppercase: true,
				IncludeLowercase: true,
				IncludeNumbers:   true,
			},
		},
		{
			name:    "Alphanumeric with symbols",
			pattern: "AS",
			expected: password.Pattern{
				IncludeUppercase: true,
				IncludeLowercase: true,
				IncludeNumbers:   true,
				IncludeSymbols:   true,
			},
		},
		{
			name:    "Custom pattern",
			pattern: "WIN",
			expected: password.Pattern{
				IncludeUppercase: true,
				IncludeLowercase: true,
				IncludeNumbers:   true,
			},
		},
		{
			name:    "Only symbols",
			pattern: "S",
			expected: password.Pattern{
				IncludeSymbols: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cmd.ParsePattern(tt.pattern)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParsePattern() = %v, want %v", result, tt.expected)
			}
		})
	}
}
