package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:',.<>?/`~"
)

func generatePassword(length int, useUppercase, useNumbers, useSpecial bool) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("password length must be at least 1")
	}

	// Build the character set based on user options
	charSet := lowercase
	if useUppercase {
		charSet += uppercase
	}
	if useNumbers {
		charSet += numbers
	}
	if useSpecial {
		charSet += specialChars
	}

	if len(charSet) == 0 {
		return "", fmt.Errorf("no character set selected for password generation")
	}

	// Generate the password
	var password strings.Builder
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", fmt.Errorf("error generating random number: %v", err)
		}
		password.WriteByte(charSet[randomIndex.Int64()])
	}

	return password.String(), nil
}

func main() {
	var length int
	var includeUppercase, includeNumbers, includeSpecial string

	fmt.Println("Random Password Generator")
	fmt.Println("=========================")

	// Get password length from the user
	fmt.Print("Enter the desired password length: ")
	_, err := fmt.Scan(&length)
	if err != nil || length < 1 {
		fmt.Println("Invalid length. Please enter a positive integer.")
		return
	}

	// Get options for including character types
	fmt.Print("Include uppercase letters? (y/n): ")
	fmt.Scan(&includeUppercase)
	fmt.Print("Include numbers? (y/n): ")
	fmt.Scan(&includeNumbers)
	fmt.Print("Include special characters? (y/n): ")
	fmt.Scan(&includeSpecial)

	// Generate the password
	password, err := generatePassword(
		length,
		strings.ToLower(includeUppercase) == "y",
		strings.ToLower(includeNumbers) == "y",
		strings.ToLower(includeSpecial) == "y",
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Output the generated password
	fmt.Printf("Generated Password: %s\n", password)
}
