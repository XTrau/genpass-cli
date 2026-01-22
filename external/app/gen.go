package app

import "math/rand"

const (
	AlphabetLowerCase string = "qwertyuiopasdfghjklzxcvbnm"
	AlphabetUpperCase string = "QWERTYUIOPASDFGHJKLZXCVBNM"
	Numbers           string = "1234567890"
)

func getCharSet(c *Config) string {
	var charSet string
	if c.nums == true {
		charSet += Numbers
	}
	if c.alphabetLower == true {
		charSet += AlphabetLowerCase
	}
	if c.alphabetUpper == true {
		charSet += AlphabetUpperCase
	}
	if len(charSet) == 0 {
		charSet += AlphabetLowerCase
	}
	return charSet
}

func GeneratePassword(c *Config) string {
	charSet := getCharSet(c)
	var result string

	for range c.size {
		index := rand.Intn(len(charSet))
		result += charSet[index : index+1]
	}

	return result
}
