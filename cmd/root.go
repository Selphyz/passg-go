package cmd

import (
	"fmt"
	"strings"

	"github.com/Selphyz/passg/pkg/password"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var (
	copyToClipboard bool
	passwordLength  int
	patternString   string
)

var rootCmd = &cobra.Command{
	Use:   "passg",
	Short: "passg is a CLI application for generating passwords",
	Long:  `passg is a command-line interface application that generates random passwords based on specified patterns.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flag("long").Changed && passwordLength <= 0 {
			return fmt.Errorf("invalid password length: %d. Please provide a positive integer", passwordLength)
		}

		length := password.DefaultLength
		if passwordLength > 0 {
			length = passwordLength
		}

		pattern := ParsePattern(patternString)
		pass := password.Generate(length, pattern)
		fmt.Printf("Generated password: %s\n", pass)

		if copyToClipboard {
			err := clipboard.WriteAll(pass)
			if err != nil {
				return fmt.Errorf("failed to copy to clipboard: %w", err)
			}
			fmt.Println("Password copied to clipboard!")
		}

		return nil
	},
}

func ParsePattern(pattern string) password.Pattern {
	p := password.Pattern{}

	if pattern == "" {
		p.IncludeUppercase = true
		p.IncludeLowercase = true
		p.IncludeNumbers = true
		return p
	}

	if strings.Contains(pattern, "A") {
		p.IncludeUppercase = true
		p.IncludeLowercase = true
		p.IncludeNumbers = true
		if strings.Contains(pattern, "S") {
			p.IncludeSymbols = true
		}
		return p
	}

	p.IncludeUppercase = strings.Contains(pattern, "W") || strings.Contains(pattern, "Y")
	p.IncludeLowercase = strings.Contains(pattern, "W") || strings.Contains(pattern, "I")
	p.IncludeNumbers = strings.Contains(pattern, "N")
	p.IncludeSymbols = strings.Contains(pattern, "S")

	return p
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVarP(&copyToClipboard, "copy", "C", false, "Copy the generated password to clipboard")
	rootCmd.Flags().IntVarP(&passwordLength, "long", "L", 0, "Specify the length of the password")
	rootCmd.Flags().StringVarP(&patternString, "pattern", "P", "", "Specify the pattern for password generation (W, Y, I, N, S, A)")
}
