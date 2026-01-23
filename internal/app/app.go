package app

import (
	"errors"
	"fmt"
)

func Run(args []string) error {
	config := getConfig()
	if config.size <= 0 {
		return errors.New("Password length must be positive number")
	}

	password := GeneratePassword(config)
	fmt.Println(password)

	return nil
}
