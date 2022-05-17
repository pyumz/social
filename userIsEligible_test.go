package main

import (
	"errors"
	"testing"
)

func TestUserIsEligible(t *testing.T) {
	var tests = []struct {
		email         string
		password      string
		age           int
		expectedError error
	}{
		{
			email:         "test@exmple.com",
			password:      "12345",
			age:           18,
			expectedError: nil,
		},
		{
			email:         "test@example.com",
			password:      "1234",
			age:           10,
			expectedError: errors.New("ge must be at least 18 years old"),
		},
	}

	for _, tt := range tests {
		err := userIsEligible(tt.email, tt.password, tt.age)
		errString := ""
		expectedErrString := ""

		if err != nil {
			errString = err.Error()
		}

		if tt.expectedError != nil {
			expectedErrString = tt.expectedError.Error()
		}

		if errString != expectedErrString {
			t.Errorf("got %s, want %s", errString, expectedErrString)
		}
	}

}
