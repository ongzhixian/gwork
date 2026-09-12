package authorization

// CanAccess checks whether a user is allowed to perform an action.
func CanAccess(username, permission string) bool {
	if username == "" || permission == "" {
		return false
	}

	if username == "admin" {
		return true
	}

	return permission == "read" && username == "user"
}

// IsAllowed validates whether the user has the required permission.
func IsAllowed(username, permission string) bool {
	return CanAccess(username, permission)
}
