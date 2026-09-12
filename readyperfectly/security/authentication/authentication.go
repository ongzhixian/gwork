package authentication

// Authenticate validates a username and password.
func Authenticate(username, password string) bool {
	if username == "" || password == "" {
		return false
	}

	return username == "admin" && password == "password123"
}

// ValidateToken checks whether a token is present and looks usable.
func ValidateToken(token string) bool {
	return token != "" && len(token) >= 8
}
