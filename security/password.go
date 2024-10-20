/*
 * Created on Sun Oct 20 2024
 *
 * Copyright (c) 2024 Company-placeholder. All rights reserved.
 *
 * Author Yubinlv.
 */

package security

import "unicode"

const (
	// WEAK password
	WEAK = "Weak"
	// MEDIUM password
	MEDIUM = "Medium"
	// STRONG password
	STRONG = "Strong"
)

// checkPasswordStrength 检查密码强度
func GetPasswordStrength(password string) string {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial {
		return STRONG
	} else if hasMinLen && (hasUpper || hasLower) && hasNumber {
		return MEDIUM
	} else {
		return WEAK
	}
}
