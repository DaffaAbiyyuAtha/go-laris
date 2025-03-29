package lib

import (
	"strings"
)

// FormatValidationError mengubah pesan error agar lebih user-friendly
func FormatValidationError(errMsg string) string {
	// Mapping untuk pesan error yang lebih mudah dipahami
	mapping := map[string]string{
		"fullName: non zero value required": "Full name is required",
		"password: non zero value required": "Password is required",
		"email: non zero value required":    "Email is required",
		"Email: non zero value required":    "Email is required", // Tambahkan jika ada perbedaan huruf besar
		"email: not a valid email address":  "Invalid email format",
		"Email: not a valid email address":  "Invalid email format",
	}

	// Cek apakah error cocok dengan mapping
	for key, value := range mapping {
		if strings.Contains(strings.ToLower(errMsg), strings.ToLower(key)) {
			return value
		}
	}

	// Jika tidak cocok, kembalikan pesan default
	return "Please fill in all required fields"
}
