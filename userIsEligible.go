package main

import (
	"errors"
	"fmt"
)

func userIsEligible(email, password string, age int) error {
	if email == "" {
		return errors.New("eMail cannot be empty")
	}

	if password == "" {
		return errors.New("password cannot be empty")
	}

	const minAge = 18
	if age < 18 {
		return fmt.Errorf("Age must be at least %v years old", minAge)
	}

	return nil
}
