package main

// At least 5 characters long but no more than 12 characters.
// Contains at least one uppercase letter.
// Contains at least one digit.

func isValidPassword(password string) bool {
	if len(password) > 12 || len(password) < 5 {
		return false
	}

	hasUppercase := false
	hasDigit := false

	for _, ch := range password {
		if ch >= 'A' && ch <= 'Z' {
			hasUppercase = true
		}

		if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}

	if !hasUppercase || !hasDigit {
		return false
	}

	return true

}
